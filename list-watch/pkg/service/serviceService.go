package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-client-practice/constant"
	resourceOps "go-client-practice/pkg/ops"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
)

type ServiceInfo struct {
	Name       string      `json:"name"`
	Status     string      `json:"status"`
	CreateTime metav1.Time `json:"createTime"`
	EndPoints  []string    `json:"endPoints"`
}

type ServiceBody struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Labels    map[string]string `json:"labels"`
}

func ListService(ctx *gin.Context) {
	serviceOps := resourceOps.ServiceOps{}
	namespace, exist := ctx.GetQuery("namespace")
	if !exist {
		ctx.JSON(constant.HttpNotFound, "wrong parameter!")
		return
	}
	list := serviceOps.List(namespace)
	if nil == list {
		ctx.JSON(constant.HttpNotFound, nil)
		return
	}
	var result []ServiceInfo
	for _, service := range list.Items {
		result = append(result, ServiceInfo{
			Status:     service.Status.String(),
			Name:       service.Name,
			CreateTime: service.CreationTimestamp,
			EndPoints:  service.Spec.ClusterIPs,
		})
	}
	ctx.JSON(constant.HttpOk, result)
}

func GetService(ctx *gin.Context) {
	serviceOps := resourceOps.ServiceOps{}
	namespace, exist := ctx.GetQuery("namespace")
	if !exist {
		ctx.JSON(constant.HttpNotFound, "wrong parameter!")
		return
	}
	name, exist := ctx.GetQuery("name")
	if !exist {
		ctx.JSON(constant.HttpOk, "wrong parameter, name is need!")
		return
	}
	service := serviceOps.Get(name, namespace)
	if nil == service {
		ctx.JSON(constant.HttpNotFound, nil)
		return
	}
	result := ServiceInfo{
		Status:     service.Status.String(),
		Name:       service.Name,
		CreateTime: service.CreationTimestamp,
		EndPoints:  service.Spec.ClusterIPs,
	}
	ctx.JSON(constant.HttpOk, result)
}

func AddService(ctx *gin.Context) {
	serviceOps := resourceOps.ServiceOps{}
	var serviceBody []byte
	serviceToAdd := ServiceBody{}
	if _, err := ctx.Request.Body.Read(serviceBody); nil != err {
		ctx.JSON(constant.HttpOk, err)
		return
	}
	if nil == serviceBody {
		ctx.JSON(constant.HttpOk, "input is empty")
		return
	}

	if err := json.Unmarshal(serviceBody, serviceToAdd); err != nil {
		ctx.JSON(constant.HttpOk, fmt.Sprintf("unmarshal input service failed, input data is "))
		return
	}

	if err := serviceOps.Add(serviceToAdd.Name, serviceToAdd.Namespace, serviceToAdd.Labels); nil != err {
		ctx.JSON(constant.HttpOk, err.Error())
	}
	ctx.JSON(constant.HttpOk, "success")
}

func DeleteService(ctx *gin.Context) {
	serviceOps := resourceOps.ServiceOps{}
	name, exist := ctx.GetQuery("name")
	if !exist {
		ctx.JSON(constant.HttpOk, "wrong parameter, name is need!")
		return
	}
	namespace, exist := ctx.GetQuery("namespace")
	if !exist {
		ctx.JSON(constant.HttpOk, "wrong parameter, namespace is need!")
		return
	}

	if err := serviceOps.Delete(name, namespace); nil != err {
		ctx.JSON(constant.HttpOk, err.Error())
	}
	ctx.JSON(constant.HttpOk, "success")
}
