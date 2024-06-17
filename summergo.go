package summergo_web

import (
	"github.com/gin-gonic/gin"
	"github.com/summer-gonner/summergo-web/controllers"

	"strconv"
)

// RegisterRoutes 注册所有路由
func registerRoutes(router *gin.Engine) {
	baseController := &controllers.BaseController{}

	// 注册认证相关路由
	baseController.AuthRoutes(router)
}
func SummerGo() {
	init, err := ApplicationInit()
	if err != nil {
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
