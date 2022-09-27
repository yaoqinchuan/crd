package ops

import (
	"fmt"
	"go-client-practice/setting"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

var DynamicClientHandler dynamic.Interface

func init() {
	config, err := clientcmd.BuildConfigFromFlags("", setting.GetConfig().GetString("cert.path"))
	if err != nil {
		panic(err.Error())
	}
	DynamicClientHandler, err = dynamic.NewForConfig(config)
	fmt.Println("init dynamic client set success, kubernetes version is")
}
