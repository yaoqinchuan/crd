package service

import (
	"github.com/gin-gonic/gin"
	"go-client-practice/constant"
	resourceOps "go-client-practice/pkg/ops"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NamespaceInfo struct {
	Name       string      `json:"name"`
	Status     string      `json:"status"`
	CreateTime metav1.Time `json:"createTime"`
}

func ListNamespace(ctx *gin.Context) {
	namespaceOps := resourceOps.NamespaceOps{}
	list := namespaceOps.List()
	if nil == list {
		ctx.JSON(constant.HttpNotFound, nil)
	}
	var result []NamespaceInfo
	for _, ns := range list.Items {
		result = append(result, NamespaceInfo{
			Status:     string(ns.Status.Phase),
			Name:       ns.Name,
			CreateTime: ns.CreationTimestamp,
		})
	}
	ctx.JSON(constant.HttpOk, result)
}

func AddNamespace(ctx *gin.Context) {
	namespaceOps := resourceOps.NamespaceOps{}
	name, exist := ctx.GetQuery("name")
	if !exist {
		ctx.JSON(constant.HttpOk, "wrong parameter, name is need!")
		return
	}
	err := namespaceOps.Add(name)
	if nil != err {
		ctx.JSON(constant.HttpOk, err.Error())
	}
	ctx.JSON(constant.HttpOk, "success")
}

func DeleteNamespace(ctx *gin.Context) {
	namespaceOps := resourceOps.NamespaceOps{}
	name, exist := ctx.GetQuery("name")
	if !exist {
		ctx.JSON(constant.HttpOk, "wrong parameter, name is need!")
		return
	}
	err := namespaceOps.Delete(name)
	if nil != err {
		ctx.JSON(constant.HttpOk, err.Error())
	}
	ctx.JSON(constant.HttpOk, "success")
}
