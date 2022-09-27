package informers

import (
	"fmt"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/client-go/informers"
	informernetworkingv1 "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/util/workqueue"
)

type IngressHandler struct {
	queue workqueue.RateLimitingInterface
}

func (p *IngressHandler) OnAdd(obj interface{}) {
	if ingress, ok := obj.(*networkingv1.Ingress); ok {
		fmt.Println("ingress ", ingress.Name, " added")
	}

}
func (p *IngressHandler) OnUpdate(oldObj, newObj interface{}) {
	if ingress, ok := newObj.(*networkingv1.Ingress); ok {
		fmt.Println("ingress ", ingress.Name, " update")
	}
}
func (p *IngressHandler) OnDelete(obj interface{}) {
	if ingress, ok := obj.(*networkingv1.Ingress); ok {
		fmt.Println("ingress ", ingress.Name, " deleted")
	}
}

func initIngressInformers(factory informers.SharedInformerFactory) *informernetworkingv1.IngressInformer {
	ingresses := factory.Networking().V1().Ingresses()
	informer := ingresses.Informer()
	informer.AddEventHandler(&IngressHandler{
		queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "ingress"),
	})
	return &ingresses
}
