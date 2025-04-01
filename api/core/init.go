package core

import (
	"api/core/rbac"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func init() {
	slog.Info("初始化核心模块....")
}

func AddRoutes(r *gin.Engine) {
	slog.Info("注册核心路由....")
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello, World!")
	})
	rbac.AddRoutes(r)

}
