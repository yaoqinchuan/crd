package ops

import (
	"context"
	"fmt"
	"google.golang.org/appengine/log"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeploymentOps struct {
}

func (ops *DeploymentOps) Add(namespace, name, imageName string) error {
	var replicas int32 = 1
	deployment := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nginx",
			Labels: map[string]string{
				"app": name,
				"env": "dev",
			},
		},
		Spec: v1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": name,
					"env": "dev",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: "nginx",
					Labels: map[string]string{
						"app": name,
						"env": "dev",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  name,
							Image: imageName, // "nginx:1.16.1"
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	if _, err := ClientSetHandler.AppsV1().Deployments(namespace).Create(context.Background(), deployment, metav1.CreateOptions{}); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (ops *DeploymentOps) List(namespace string) *v1.DeploymentList {
	deploys, err := ClientSetHandler.AppsV1().Deployments(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Errorf(context.Background(), "list deployment failed, error %v", err)
		return nil
	}
	return deploys
}

func (ops *DeploymentOps) Get(name, namespace string) *v1.Deployment {
	deploys, err := ClientSetHandler.AppsV1().Deployments(namespace).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		log.Errorf(context.Background(), "get deployment failed, error %v", err)
		return nil
	}
	return deploys
}

func (ops *DeploymentOps) Delete(name, namespace string) error {
	err := ClientSetHandler.AppsV1().Deployments(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
