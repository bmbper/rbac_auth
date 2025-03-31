package main

import (
	_ "api/core"

	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化Gin路由
	r := gin.Default()

	r.Run(":36000")
}
