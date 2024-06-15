package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/summer-gonner/summergo-web/services"
)

func (uc *BaseController) AuthRoutes(router *gin.Engine) {
	authController := &services.AuthController{}
	// 注册用户相关的路由
	router.GET("/users/:id", authController.GetUser)
	router.POST("/users", authController.CreateUser)
	router.PUT("/users/id", authController.UpdateUser)
	router.DELETE("/users/id", authController.DeleteUser)
}
