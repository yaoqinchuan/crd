package main

import (
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
)

func main() {
	configPath := "config"
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if nil != err {
		panic(err)
	}
	clientSet, err := kubernetes.NewForConfig(config)
	//factory := informers.NewSharedInformerFactory(clientSet, 0)
	//指定namespace
	factory := informers.NewSharedInformerFactoryWithOptions(clientSet, 0, informers.WithNamespace("default"))
	informer := factory.Core().V1().Pods().Informer()

	queue := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "controller")
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			println("add event")
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if nil != err {
				panic(err)
			}
			queue.AddRateLimited(key)
		},
		DeleteFunc: func(obj interface{}) {
			println("delete event")
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if nil != err {
				panic(err)
			}
			queue.AddRateLimited(key)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			println("update event")
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			if nil != err {
				panic(err)
			}
			queue.AddRateLimited(key)
		},
	})
	stopCh := make(chan struct{})
	factory.Start(stopCh)
	factory.WaitForCacheSync(stopCh)
	<-stopCh

}
