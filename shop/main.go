package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_shop/shop/routers"
	"go_shop/shop/util"
	"go_shop/tools/logger"
)

func main() {
	config, err := util.GetConfig()
	if err != nil {
		logger.Error.Println(err)
		panic("获取配置文件异常")
	}

	r := gin.Default()
	routers.Entry(r)
	err = r.Run(fmt.Sprint(":", config.Server.Port))
	if err != nil {
		panic(err)
	}
}
