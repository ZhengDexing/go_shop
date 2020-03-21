package routers

import (
	"github.com/gin-gonic/gin"
	"go_shop/controller"
)

// 路由
func Entry(e *gin.Engine) {
	e.POST("/registered", controller.Registered)
	e.POST("/login")

	userRouter := e.Group("/user")
	userRouter.POST("/registered")
}
