package main

import (
	"github.com/gin-gonic/gin"
	"go_shop/routers"
)

func main() {
	r := gin.Default()
	routers.Entry(r)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
