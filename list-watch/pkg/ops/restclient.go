package ops

import (
	"fmt"
	"go-client-practice/setting"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var RestClientHandler *rest.RESTClient

func init() {
	config, err := clientcmd.BuildConfigFromFlags("", setting.GetConfig().GetString("cert.path"))
	if err != nil {
		panic(err.Error())
	}
	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	RestClientHandler, err = rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}
	fmt.Println("init client set success, kubernetes version is", RestClientHandler.APIVersion())
}
