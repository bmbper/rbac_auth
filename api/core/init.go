package core

import (
	"bmbp/core/rbac"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func InitCore(r *gin.Engine) {
	slog.Info("初始化核心模块....")

	// 初始化数据库
	initCoreDb()

	// 注册核路由
	initCoreRoute(r)

}

func initCoreRoute(r *gin.Engine) {
	// 注册核心路由
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello, World!")
	})
	rbac.InitRbacRoute(r)
}

func initCoreDb() {
	rbac.InitRbacDb()
}
