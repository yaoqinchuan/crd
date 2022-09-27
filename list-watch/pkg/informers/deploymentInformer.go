package informers

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/informers"
	v1 "k8s.io/client-go/informers/apps/v1"
)

type DeploymentHandler struct {
}

func (d *DeploymentHandler) OnAdd(obj interface{}) {
	if dep, ok := obj.(*appsv1.Deployment); ok {
		fmt.Println("deployment ", dep.Name, " added")
	}
}
func (d *DeploymentHandler) OnUpdate(oldObj, newObj interface{}) {
	if dep, ok := newObj.(*appsv1.Deployment); ok {
		fmt.Println("deployment ", dep.Name, " update")
	}
}
func (d *DeploymentHandler) OnDelete(obj interface{}) {
	if dep, ok := obj.(*appsv1.Deployment); ok {
		fmt.Println("deployment ", dep.Name, " delete")
	}
}

func initDeploymentInformers(factory informers.SharedInformerFactory) *v1.DeploymentInformer {
	deployments := factory.Apps().V1().Deployments()
	informer := deployments.Informer()
	informer.AddEventHandler(&DeploymentHandler{})
	return &deployments
}
