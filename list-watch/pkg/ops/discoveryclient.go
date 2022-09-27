package ops

import (
	"fmt"
	"go-client-practice/setting"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery/cached/disk"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

var CachedDiscoveryClient *disk.CachedDiscoveryClient

func init() {
	config, err := clientcmd.BuildConfigFromFlags("", setting.GetConfig().GetString("cert.path"))
	if err != nil {
		panic(err.Error())
	}
	CachedDiscoveryClient, err = disk.NewCachedDiscoveryClientForConfig(config, "./cache/discovery", "./cache/discovery", time.Minute*60)
	if nil != err {
		panic(err.Error())
	}
	_, apiResourceList, err := CachedDiscoveryClient.ServerGroupsAndResources()
	if nil != err {
		panic(err.Error())
	}
	for _, list := range apiResourceList {
		gv, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			panic(err)
		}

		for _, resource := range list.APIResources {
			fmt.Printf("name %v, group %v, version: %v\n", resource.Name, gv.Group, gv.Version)
		}
	}
}
