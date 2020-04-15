package filter

import (
	"github.com/gin-gonic/gin"
	"go_shop/shop/util"
	"net/http"
)

// 校验是否登录和请求是否合法
func CheckLogin(context *gin.Context) {
	sign, err := context.Cookie("sign")
	if err != nil {
		context.Abort()
		context.JSON(http.StatusForbidden, util.RespOut(util.Forbidden, ""))
		return
	}
	jwt := new(util.Jwt)
	user, err := jwt.ParseToken(sign)
	if err != nil {
		context.Abort()
		context.JSON(http.StatusForbidden, util.RespOut(util.Forbidden, ""))
		return
	}
	context.Set("userId", user.Id)
	context.Set("nickName", user.NickName)
	context.Next()
}
