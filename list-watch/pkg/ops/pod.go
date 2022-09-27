package ops

import (
	"context"
	"fmt"
	"google.golang.org/appengine/log"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type PodOps struct {
}

func (ops *PodOps) List(namespace string) *v1.PodList {
	resource := schema.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}
	podsUnstruct, err := DynamicClientHandler.Resource(resource).Namespace(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	podList := &v1.PodList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(podsUnstruct.UnstructuredContent(), podList)
	if err != nil {
		panic(err)
	}
	return podList
}

func (ops *PodOps) Get(name, namespace string) *v1.Pod {
	pod, err := ClientSetHandler.CoreV1().Pods(namespace).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		log.Errorf(context.Background(), "list namespace failed, error %v", err)
		return nil
	}
	return pod
}

func (ops *PodOps) Delete(name, namespace string) error {
	err := ClientSetHandler.CoreV1().Pods(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
