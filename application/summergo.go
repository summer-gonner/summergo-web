package application

import (
	"github.com/gin-gonic/gin"
	"github.com/summer-gonner/summergo-web/controllers"

	"github.com/summer-gonner/summergo-web/pkg"
	"strconv"
)

// RegisterRoutes 注册所有路由
func registerRoutes(router *gin.Engine) {
	baseController := &controllers.BaseController{}

	// 注册认证相关路由
	baseController.AuthRoutes(router)
}
func SummerGo() {
	init, error := pkg.ApplicationInit()
	if error != nil {
		return
	}

	r := gin.Default()
	registerRoutes(r)
	if init != nil {
		add := ":" + strconv.Itoa(init.Server.Port)
		err := r.Run(add)
		if err != nil {
			return
		}
	} else {
		add := ":8080"
		err := r.Run(add)
		if err != nil {
			return
		}
	}

}
