package main

import (
	"fmt"
	"github.com/yaoqinchuan/operator-demo/ingress/pkg"
	"k8s.io/client-go/informers"
	_ "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	_ "k8s.io/client-go/tools/clientcmd"
	_ "k8s.io/client-go/util/workqueue"
)

func main() {
	configPath := "config"
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if nil != err {
		panic(err)
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if nil != err {
		panic(err)
	}
	factory := informers.NewSharedInformerFactory(clientSet, 0)

	serviceInformer := factory.Core().V1().Services()

	ingressInformer := factory.Networking().V1().Ingresses()
	controller := pkg.NewController(clientSet, serviceInformer, ingressInformer)
	stopChan := make(chan struct{})
	fmt.Println("启动infromer中。。。。")
	factory.Start(stopChan)
	factory.WaitForCacheSync(stopChan)
	controller.Run(stopChan)
	<-stopChan

}
