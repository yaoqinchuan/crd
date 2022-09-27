package ops

import (
	"context"
	"fmt"
	"google.golang.org/appengine/log"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NamespaceOps struct {
}

func (ops *NamespaceOps) List() *v1.NamespaceList {
	ns, err := ClientSetHandler.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Errorf(context.Background(), "list namespace failed, error %v", err)
		return nil
	}
	return ns
}
func (ops *NamespaceOps) Add(name string) error {
	namespace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	_, err := ClientSetHandler.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (ops *NamespaceOps) Delete(name string) error {
	err := ClientSetHandler.CoreV1().Namespaces().Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (ops *NamespaceOps) CheckIfExist(name string) bool {

	_, err := ClientSetHandler.CoreV1().Namespaces().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return false
	}
	return true
}
