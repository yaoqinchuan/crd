package ops

import (
	"fmt"
	"go-client-practice/setting"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var ClientSetHandler *kubernetes.Clientset

func init() {
	config, err := clientcmd.BuildConfigFromFlags("", setting.GetConfig().GetString("cert.path"))
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	ClientSetHandler, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	version, err := ClientSetHandler.ServerVersion()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("init client set success, kubernetes version is", version)
}
