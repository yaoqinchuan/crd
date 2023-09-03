package utils

import (
	"bytes"
	yqctechv1alpha1 "github.com/api/v1alpha1"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"text/template"
)

func parseTemplate(templateName string, controller *yqctechv1alpha1.ServiceIngressController) []byte {
	tmpl, err := template.ParseFiles("internal/controller/template/" + templateName + ".yml")
	if err != nil {
		panic(err)
	}
	b := new(bytes.Buffer)
	err = tmpl.Execute(b, controller)
	if err != nil {
		panic(err)
	}
	return b.Bytes()
}

func NewDeployment(controller *yqctechv1alpha1.ServiceIngressController) *v1.Deployment {
	deployment := &v1.Deployment{}
	err := yaml.Unmarshal(parseTemplate("deployment", controller), deployment)
	if nil != err {
		panic(err)
	}
	return deployment
}

func NewService(controller *yqctechv1alpha1.ServiceIngressController) *corev1.Service {
	service := &corev1.Service{}
	err := yaml.Unmarshal(parseTemplate("service", controller), service)
	if nil != err {
		panic(err)
	}
	return service
}

func NewIngress(controller *yqctechv1alpha1.ServiceIngressController) *netv1.Ingress {
	service := &netv1.Ingress{}
	err := yaml.Unmarshal(parseTemplate("ingress", controller), service)
	if nil != err {
		panic(err)
	}
	return service
}
