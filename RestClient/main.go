package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	_ "k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	print("hello\n")
	configPath := "config"
	// config
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		panic(err)
	}
	// RestClient
	/*
		config.GroupVersion = &corev1.SchemeGroupVersion
		config.NegotiatedSerializer = scheme.Codecs
		config.APIPath = "api"
		RestClient, err := rest.RESTClientFor(config)
		if err != nil {
			panic(err)
		}
		// data
		pod := corev1.Pod{}
		err = RestClient.Get().Namespace("kube-system").Resource("pods").Name("calico-node-dhpnm").Do(context.TODO()).Into(&pod)
		if err != nil {
			panic(err)
		} else {
			print(pod.Name)
		}
	*/

	// clientset
	/*
		clientset, err := kubernetes.NewForConfig(config)
		if nil != err {
			panic(err)
		}
		podClients := clientset.CoreV1().Pods("kube-system")
		list, err := podClients.List(context.TODO(), metaV1.ListOptions{})
		if nil != err {
			panic(err)
		}
		for _, pod := range list.Items {
			fmt.Printf("NameSpace:%v \t Name:%v \t Status:%+v\n", pod.Name, pod.Namespace, pod.Status.Phase)

		}
	*/
	// dynamicClient
	/*
		dynamicClient, err := dynamic.NewForConfig(config)
		if err != nil {
			panic(err)
		}
		gvr := scheme.GroupVersionResource{Version: "v1", Resource: "pods"}
		list, err := dynamicClient.Resource(gvr).Namespace("kube-system").List(context.TODO(), metav1.ListOptions{Limit: 100})
		if err != nil {
			panic(err.Error())
		}
		podList := &apiv1.PodList{}

		// 转换
		err = runtime.DefaultUnstructuredConverter.FromUnstructured(list.UnstructuredContent(), podList)

		if err != nil {
			panic(err.Error())
		}

		// 表头
		fmt.Printf("namespace\t status\t\t name\n")

		// 每个pod都打印namespace、status.Phase、name三个字段
		for _, d := range podList.Items {
			fmt.Printf("%v\t %v\t %v\n",
				d.Namespace,
				d.Status.Phase,
				d.Name)
		}
	*/

	// 新建discoveryClient实例

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)

	if err != nil {
		panic(err.Error())
	}

	// 获取所有分组和资源数据
	APIGroup, APIResourceListSlice, err := discoveryClient.ServerGroupsAndResources()

	if err != nil {
		panic(err.Error())
	}

	// 先看Group信息
	fmt.Printf("APIGroup :\n\n %v\n\n\n\n", APIGroup)

	// APIResourceListSlice是个切片，里面的每个元素代表一个GroupVersion及其资源
	for _, singleAPIResourceList := range APIResourceListSlice {

		// GroupVersion是个字符串，例如"apps/v1"
		groupVerionStr := singleAPIResourceList.GroupVersion

		// ParseGroupVersion方法将字符串转成数据结构
		gv, err := schema.ParseGroupVersion(groupVerionStr)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println("*****************************************************************")
		fmt.Printf("GV string [%v]\nGV struct [%#v]\nresources :\n\n", groupVerionStr, gv)

		// APIResources字段是个切片，里面是当前GroupVersion下的所有资源
		for _, singleAPIResource := range singleAPIResourceList.APIResources {
			fmt.Printf("%v\n", singleAPIResource.Name)
		}
	}

}
