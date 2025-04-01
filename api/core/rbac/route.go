package rbac

import (
	"api/core/rbac/organ"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.Engine) {
	slog.Info("注册核心-rbac模块路由")
	// 注册核心路由
	rbac := r.Group("/rbac")
	organ.AddRoutes(rbac)
}
