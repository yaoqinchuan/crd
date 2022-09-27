package service

import (
	"github.com/gin-gonic/gin"
	"go-client-practice/constant"
	resourceOps "go-client-practice/pkg/ops"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodInfo struct {
	Name       string      `json:"name"`
	Status     string      `json:"status"`
	CreateTime metav1.Time `json:"createTime"`
	Ip         string      `json:"ip"`
	Node       string      `json:"node"`
}

func ListPod(ctx *gin.Context) {
	podOps := resourceOps.PodOps{}
	namespace, exist := ctx.GetQuery("namespace")
	if !exist {
		ctx.JSON(constant.HttpNotFound, "wrong parameter!")
		return
	}
	list := podOps.List(namespace)
	if nil == list {
		ctx.JSON(constant.HttpNotFound, nil)
	}
	var result []PodInfo
	for _, pod := range list.Items {
		result = append(result, PodInfo{
			Status:     string(pod.Status.Phase),
			Name:       pod.Name,
			CreateTime: pod.CreationTimestamp,
			Ip:         pod.Status.PodIP,
			Node:       pod.Spec.NodeName,
		})
	}
	ctx.JSON(constant.HttpOk, result)
}

func GetPod(ctx *gin.Context) {
	podOps := resourceOps.PodOps{}
	namespace, exist := ctx.GetQuery("namespace")
	if !exist {
		ctx.JSON(constant.HttpNotFound, "wrong parameter!")
		return
	}
	name, exist := ctx.GetQuery("name")
	if !exist {
		ctx.JSON(constant.HttpNotFound, "wrong parameter!")
		return
	}
	pod := podOps.Get(name, namespace)
	if nil == pod {
		ctx.JSON(constant.HttpNotFound, nil)
	}
	result := PodInfo{
		Status:     string(pod.Status.Phase),
		Name:       pod.Name,
		CreateTime: pod.CreationTimestamp,
		Ip:         pod.Status.PodIP,
		Node:       pod.Spec.NodeName,
	}

	ctx.JSON(constant.HttpOk, result)
}

func DeletePod(ctx *gin.Context) {
	podOps := resourceOps.ServiceOps{}
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

	if err := podOps.Delete(name, namespace); nil != err {
		ctx.JSON(constant.HttpOk, err.Error())
	}
	ctx.JSON(constant.HttpOk, "success")
}
