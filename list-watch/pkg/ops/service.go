package ops

import (
	"context"
	"fmt"
	"google.golang.org/appengine/log"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

type ServiceOps struct {
}

func (ops *ServiceOps) List(namespace string) *v1.ServiceList {
	serviceList := &v1.ServiceList{}
	err := RestClientHandler.
		Get().
		Namespace(namespace).
		Resource("services").
		VersionedParams(&metav1.ListOptions{}, scheme.ParameterCodec).
		Do(context.Background()).Into(serviceList)
	if err != nil {
		panic(err.Error())
	}
	return serviceList
}
func (ops *ServiceOps) Get(name, namespace string) *v1.Service {
	service, err := ClientSetHandler.CoreV1().Services(namespace).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		log.Errorf(context.Background(), "list namespace failed, error %v", err)
		return nil
	}
	return service
}

func (ops *ServiceOps) Add(name, namespace string, labels map[string]string) error {
	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: v1.ServiceSpec{
			Selector: labels,
		},
	}
	_, err := ClientSetHandler.CoreV1().Services(namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (ops *ServiceOps) Delete(name, namespace string) error {
	err := ClientSetHandler.CoreV1().Services(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
