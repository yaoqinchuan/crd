package pkg

import (
	"context"
	"fmt"
	v15 "k8s.io/api/core/v1"
	v13 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	v14 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	informer "k8s.io/client-go/informers/core/v1"
	netInformer "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"
	v12 "k8s.io/client-go/listers/core/v1"
	v1 "k8s.io/client-go/listers/networking/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"reflect"
	"time"
)

const workNum = 5
const maxRetry = 10

type controller struct {
	client        kubernetes.Interface
	ingressLister v1.IngressLister
	serviceLister v12.ServiceLister
	queue         workqueue.RateLimitingInterface
}

func (c *controller) updateService(oldObj interface{}, newObj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(newObj)
	if nil != err {
		runtime.HandleError(err)
	}
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if nil != err {
		runtime.HandleError(err)
	}
	fmt.Printf("update service:%s:%s\n ", namespace, name)
	if reflect.DeepEqual(oldObj, newObj) {
		return
	}
	c.enqueue(newObj)
}

func (c *controller) addService(object interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(object)
	if nil != err {
		runtime.HandleError(err)
	}
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if nil != err {
		runtime.HandleError(err)
	}
	fmt.Printf("add service:%s:%s\n ", namespace, name)
	c.enqueue(object)
}

// 将service入队
func (c *controller) enqueue(object interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(object)
	if nil != err {
		runtime.HandleError(err)
	}
	c.queue.Add(key)
}

func (c *controller) deleteIngress(obj interface{}) {
	ingress := obj.(*v13.Ingress)
	ownerReference := v14.GetControllerOf(ingress)
	if nil == ownerReference {
		return
	}
	if "service" != ownerReference.Kind {
		return
	}
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if nil != err {
		runtime.HandleError(err)
	}
	c.queue.Add(key)
}

func (c *controller) Run(stopCh chan struct{}) {
	for i := 0; i < workNum; i++ {
		go wait.Until(c.worker, time.Minute, stopCh)
	}
}

func (c *controller) worker() {
	for c.processNextItem() {

	}
}

func (c *controller) processNextItem() bool {
	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	defer c.queue.Done(item)
	key := item.(string)
	err := c.syncService(key)
	if nil != err {
		runtime.HandleError(err)
	}
	return true
}

func (c *controller) handlerError(key string, err error) {
	if c.queue.NumRequeues(key) <= maxRetry {
		c.queue.AddRateLimited(key)
		return
	}
	runtime.HandleError(err)
	c.queue.Forget(key)
}

func (c *controller) syncService(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if nil != err {
		fmt.Printf("error occurs %s/n", err)
		return err
	}
	service, err := c.serviceLister.Services(namespace).Get(name)
	if err != nil {
		fmt.Printf("error occurs %s/n", err)
	}
	annotation := service.GetAnnotations()["ingress/http"]
	ingress, err := c.ingressLister.Ingresses(namespace).Get(name)
	if nil != err && !errors.IsNotFound(err) {
		fmt.Printf("error occurs %s/n", err)
		return err
	}
	if annotation != "" && errors.IsNotFound(err) {
		constructIngress := c.constructIngress(*service)
		_, err := c.client.NetworkingV1().Ingresses(namespace).Create(context.TODO(), constructIngress, v14.CreateOptions{})
		if err != nil {
			fmt.Printf("error occurs %s/n", err)
		}
	} else if annotation == "" && ingress != nil {
		err := c.client.NetworkingV1().Ingresses(namespace).Delete(context.TODO(), name, v14.DeleteOptions{})
		if err != nil {
			fmt.Printf("error occurs %s/n", err)
		}
	}
	return nil
}

func (c *controller) constructIngress(service v15.Service) *v13.Ingress {
	ingress := &v13.Ingress{}
	ingress.ObjectMeta.OwnerReferences = []v14.OwnerReference{*v14.NewControllerRef(&service, v13.SchemeGroupVersion.WithKind("service"))}
	ingress.Name = service.Name
	ingress.Namespace = service.Namespace
	pathType := v13.PathTypePrefix
	icn := "nginx"
	ingress.Spec = v13.IngressSpec{
		IngressClassName: &icn,
		Rules: []v13.IngressRule{
			{Host: "yqc.com",
				IngressRuleValue: v13.IngressRuleValue{
					HTTP: &v13.HTTPIngressRuleValue{
						Paths: []v13.HTTPIngressPath{{
							Path:     "/",
							PathType: &pathType,
							Backend: v13.IngressBackend{
								Service: &v13.IngressServiceBackend{
									Name: service.Name,
									Port: v13.ServiceBackendPort{Number: 80},
								},
							},
						}},
					},
				}},
		},
	}
	return ingress
}

func NewController(client kubernetes.Interface, serviceInformer informer.ServiceInformer, ingressInformer netInformer.IngressInformer) controller {
	c := controller{
		client:        client,
		ingressLister: ingressInformer.Lister(),
		serviceLister: serviceInformer.Lister(),
		queue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "ingressManager"),
	}
	serviceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.addService,
		UpdateFunc: c.updateService,
	})
	ingressInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: c.deleteIngress,
	})
	return c
}
