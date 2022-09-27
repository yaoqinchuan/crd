package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-client-practice/constant"
	resourceOps "go-client-practice/pkg/ops"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeploymentInfo struct {
	Name        string      `json:"name"`
	CreateTime  metav1.Time `json:"createTime"`
	ExpectCount int32       `json:"expectCount"`
	ReadyCount  int32       `json:"readyCount"`
}

func AddDeployment(ctx *gin.Context) {
	deploymentOps := resourceOps.DeploymentOps{}
	namespace, exist := ctx.GetQuery("namespace")
	if !exist {
		ctx.JSON(constant.HttpOk, "wrong parameter, namespace is need!")
		return
	}
	name, exist := ctx.GetQuery("name")
	if !exist {
		ctx.JSON(constant.HttpOk, "wrong parameter, name is need!")
		return
	}
	imageName, exist := ctx.GetQuery("imageName")
	if !exist {
		ctx.JSON(constant.HttpOk, "wrong parameter, image name is need!")
		return
	}
	err := deploymentOps.Add(namespace, name, imageName)
	if err != nil {
		ctx.JSON(constant.HttpInternalError, fmt.Sprintf("create deployment failed, err %v !", err.Error()))
		return
	}
}

func DeleteDeployment(ctx *gin.Context) {
	deploymentOps := resourceOps.DeploymentOps{}
	namespace, exist := ctx.GetQuery("namespace")
	if !exist {
		ctx.JSON(constant.HttpOk, "wrong parameter, namespace is need!")
		return
	}
	name, exist := ctx.GetQuery("name")
	if !exist {
		ctx.JSON(constant.HttpOk, "wrong parameter, name is need!")
		return
	}
	err := deploymentOps.Delete(name, namespace)
	if err != nil {
		ctx.JSON(constant.HttpInternalError, fmt.Sprintf("delete deployment failed, err %v !", err.Error()))
		return
	}
}

func ListDeployment(ctx *gin.Context) {
	deploymentOps := resourceOps.DeploymentOps{}
	namespace, exist := ctx.GetQuery("namespace")
	if !exist {
		ctx.JSON(constant.HttpNotFound, "wrong parameter!")
		return
	}
	list := deploymentOps.List(namespace)
	if nil == list {
		ctx.JSON(constant.HttpNotFound, nil)
		return
	}
	var result []DeploymentInfo
	for _, deploy := range list.Items {
		result = append(result, DeploymentInfo{
			Name:        deploy.Name,
			CreateTime:  deploy.CreationTimestamp,
			ExpectCount: deploy.Status.Replicas,
			ReadyCount:  deploy.Status.ReadyReplicas,
		})
	}
	ctx.JSON(constant.HttpOk, result)
}
func GetDeployment(ctx *gin.Context) {
	deploymentOps := resourceOps.DeploymentOps{}
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
	deploy := deploymentOps.Get(name, namespace)
	if nil == deploy {
		ctx.JSON(constant.HttpNotFound, nil)
		return
	}
	result := DeploymentInfo{
		Name:        deploy.Name,
		CreateTime:  deploy.CreationTimestamp,
		ExpectCount: deploy.Status.Replicas,
		ReadyCount:  deploy.Status.ReadyReplicas,
	}
	ctx.JSON(constant.HttpOk, result)
}
