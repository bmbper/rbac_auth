package main

import (
	"bmbp/base"
	"bmbp/core"
	"bmbp/db"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func main() {
	slog.Info("应用启动....")
	slog.Info("加载配置文件...")
	// 加载配置文件
	err := base.LoadAppConfig("./config.yaml")
	if err != nil {
		slog.Error("应用启动失败：无效的配置文件->", "error", err)
		return
	}

	// 初始化数据库
	err = db.InitDB()
	if err != nil {
		slog.Error("应用启动失败：数据库初始化失败->", "error", err)
		return
	}

	// 初始化Gin路由
	r := gin.Default()
	r.Use(gin.Recovery())

	// 注册核心路由
	core.InitCore(r)

	addr := fmt.Sprintf("%s:%d", base.GlobalConfig.Server.Host, base.GlobalConfig.Server.Port)
	r.Run(addr)
}
