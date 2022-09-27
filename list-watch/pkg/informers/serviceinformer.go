package informers

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	v1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/util/workqueue"
)

type ServiceHandler struct {
	queue workqueue.RateLimitingInterface
}

func (p *ServiceHandler) OnAdd(obj interface{}) {
	if service, ok := obj.(*corev1.Service); ok {
		fmt.Println("service ", service.Name, " added")
	}

}
func (p *ServiceHandler) OnUpdate(oldObj, newObj interface{}) {
	if service, ok := newObj.(*corev1.Service); ok {
		fmt.Println("service ", service.Name, " update")
	}
}
func (p *ServiceHandler) OnDelete(obj interface{}) {
	if service, ok := obj.(*corev1.Service); ok {
		fmt.Println("service ", service.Name, " deleted")
	}
}

func initServiceInformers(factory informers.SharedInformerFactory) *v1.ServiceInformer {
	services := factory.Core().V1().Services()
	informer := services.Informer()
	informer.AddEventHandler(&ServiceHandler{
		queue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "service"),
	})
	return &services
}
