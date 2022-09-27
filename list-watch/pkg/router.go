package pkg

import (
	"github.com/gin-gonic/gin"
	"go-client-practice/pkg/service"
)

func SetRouter(r *gin.Engine) {
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, "hello")
	})
	v1 := r.Group("/v1")
	v1.GET("/namespaces", service.ListNamespace)
	v1.POST("/namespaces", service.AddNamespace)
	v1.DELETE("/namespaces", service.DeleteNamespace)

	v1.GET("/deployments", service.ListDeployment)
	v1.GET("/deployment", service.GetDeployment)
	v1.POST("/deployment", service.AddDeployment)
	v1.DELETE("/deployment", service.DeleteDeployment)

	v1.GET("/services", service.ListService)
	v1.GET("/service", service.GetService)
	v1.POST("/service", service.AddService)
	v1.DELETE("/service", service.DeleteService)

	v1.GET("/pods", service.ListPod)
	v1.GET("/pod", service.GetPod)
	v1.DELETE("/pod", service.DeletePod)
}
