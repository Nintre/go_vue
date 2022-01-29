package main

import (
	"github.com/gin-gonic/gin"
	"hutaiyi.study/gin_vue/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.GET("/api/auth/register", controller.Register)
	return r
}