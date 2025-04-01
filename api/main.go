package main

import (
	"api/core"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func main() {
	slog.Info("应用启动")
	// 初始化Gin路由
	r := gin.Default()
	r.Use(gin.Recovery())
	// 注册核心路由
	core.AddRoutes(r)
	r.Run(":36000")
}
