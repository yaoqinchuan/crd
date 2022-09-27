package main

import (
	"github.com/gin-gonic/gin"
	"go-client-practice/pkg"
	"go-client-practice/pkg/informers"
	_ "go-client-practice/pkg/ops"
)

func main() {
	r := gin.Default()
	pkg.SetRouter(r)
	stopCh := make(chan struct{})
	go informers.InitInformers(stopCh)
	err := r.Run()
	if err != nil {
		return
	}
}
