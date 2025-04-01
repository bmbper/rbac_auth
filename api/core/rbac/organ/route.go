package organ

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.RouterGroup) {
	slog.Info("注册核心-rbac-组织模块路由")
	organ := r.Group("/organ")
	{
		organ.POST("/tree", OrganTree)
		organ.POST("/page", OrganPage)
		organ.POST("/list", OrganList)
		organ.POST("/info", OrganInfo)
		organ.POST("/save", OrganSave)
		organ.POST("/enable", OrganEnable)
		organ.POST("/disable", OrganDisable)
		organ.POST("/delete", OrganDel)
	}
}
