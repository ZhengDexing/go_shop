package routers

import (
	"github.com/gin-gonic/gin"
	"go_shop/shop/controller"
	"go_shop/shop/filter"
	"go_shop/shop/security"
)

// 路由
func Entry(e *gin.Engine) {
	e.Use(filter.RoutersLogger)
	e.POST("/registered", controller.Registered)
	e.POST("/login")

	// 以下接口需要校验请求合法性
	e.Use(security.CheckLogin)
	userRouter := e.Group("/user")
	userRouter.POST("/registered")
}
