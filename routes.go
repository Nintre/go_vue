package main

import (
	"hutaiyi.study/gin_vue/middleware"

	"github.com/gin-gonic/gin"
	"hutaiyi.study/gin_vue/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}
