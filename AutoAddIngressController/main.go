package main

import (
	_ "k8s.io/client-go/informers"
	_ "k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/tools/cache"
	_ "k8s.io/client-go/tools/clientcmd"
	_ "k8s.io/client-go/util/workqueue"
)

func main() {

}
