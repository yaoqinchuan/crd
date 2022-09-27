package informers

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	v1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type PodHandler struct {
	queue workqueue.RateLimitingInterface
}

func (p *PodHandler) OnAdd(obj interface{}) {
	if pod, ok := obj.(*corev1.Pod); ok {
		fmt.Println("pod ", pod.Name, " added")
		key, err := cache.MetaNamespaceKeyFunc(obj)
		if err != nil {
			panic(fmt.Sprintf("can not get key for error, %v", err.Error()))
		}
		p.queue.AddRateLimited(key)
	}

}
func (p *PodHandler) OnUpdate(oldObj, newObj interface{}) {
	if pod, ok := newObj.(*corev1.Pod); ok {
		fmt.Println("pod ", pod.Name, " update")
		key, err := cache.MetaNamespaceKeyFunc(newObj)
		if err != nil {
			panic(fmt.Sprintf("can not get key for error, %v", err.Error()))
		}
		p.queue.AddRateLimited(key)
	}
}
func (p *PodHandler) OnDelete(obj interface{}) {
	if pod, ok := obj.(*corev1.Pod); ok {
		fmt.Println("pod ", pod.Name, " deleted")
		key, err := cache.MetaNamespaceKeyFunc(obj)
		if err != nil {
			panic(fmt.Sprintf("can not get key for error, %v", err.Error()))
		}
		p.queue.AddRateLimited(key)
	}
}

func initPodInformers(factory informers.SharedInformerFactory) *v1.PodInformer {
	pods := factory.Core().V1().Pods()
	informer := pods.Informer()
	informer.AddEventHandler(&PodHandler{
		queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "controller"),
	})
	return &pods
}
