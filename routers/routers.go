package routers

import (
	"github.com/gin-gonic/gin"
	"go_shop/controller"
	"go_shop/security"
	"go_shop/util"
)

// 路由
func Entry(e *gin.Engine) {
	e.Use(util.RoutersLogger)
	e.POST("/registered", controller.Registered)
	e.POST("/login")

	// 以下接口需要校验请求合法性
	e.Use(security.CheckLogin)
	userRouter := e.Group("/user")
	userRouter.POST("/registered")
}
