package security

import (
	"github.com/gin-gonic/gin"
	"go_shop/util"
	"net/http"
)

// 校验是否登录和请求是否合法
func CheckLogin(context *gin.Context) {
	sign, err := context.Cookie("sign")
	if err != nil {
		context.JSON(http.StatusForbidden, util.RespOut(util.Forbidden, ""))
		return
	}
	jwt := new(Jwt)
	_, err = jwt.ParseToken(sign)
	if err != nil {
		context.JSON(http.StatusForbidden, util.RespOut(util.Forbidden, ""))
		return
	}
	context.Next()
}
