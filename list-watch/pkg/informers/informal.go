package informers

import (
	"go-client-practice/pkg/controller"
	"go-client-practice/pkg/ops"
	"k8s.io/client-go/informers"
)

func InitInformers(stopCh chan struct{}) {
	factory := informers.NewSharedInformerFactoryWithOptions(ops.ClientSetHandler, 0, informers.WithNamespace("yqc"))
	initDeploymentInformers(factory)
	initPodInformers(factory)
	serviceInformer := initServiceInformers(factory)
	ingressInformer := initIngressInformers(factory)
	initDeploymentInformers(factory)
	factory.Start(stopCh)
	factory.WaitForCacheSync(stopCh)
	newController := controller.NewController(ops.ClientSetHandler, *serviceInformer, *ingressInformer)
	newController.Run(stopCh)
	<-stopCh
}
