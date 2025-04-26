package rbac

import (
	"bmbp/core/rbac/organ"
	"bmbp/db"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func InitRbacRoute(r *gin.Engine) {
	slog.Info("注册Core-rbac模块路由")
	// 注册核心路由
	rbac := r.Group("/rbac")
	initRbacOrganRoute(rbac)
}
func InitRbacDb() {
	initRbacOrganDb()
}
func initRbacOrganRoute(r *gin.RouterGroup) {
	slog.Info("注册核心-rbac-组织模块路由")
	organ_group := r.Group("/v1/organ")
	{
		organ_group.POST("/tree", organ.OrganTree)
		organ_group.POST("/page", organ.OrganPage)
		organ_group.POST("/list", organ.OrganList)
		organ_group.POST("/info", organ.OrganInfo)
		organ_group.POST("/save", organ.OrganSave)
		organ_group.POST("/enable", organ.OrganEnable)
		organ_group.POST("/disable", organ.OrganDisable)
		organ_group.POST("/delete", organ.OrganDel)
	}
}
func initRbacOrganDb() {
	db.DbUtil.AutoMigrate(&organ.BmbpRbacOrgan{})
	db.DbUtil.AutoMigrate(&organ.BmbpRbacOrganArea{})
	db.DbUtil.AutoMigrate(&organ.BmbpRbacOrganGroup{})
	db.DbUtil.AutoMigrate(&organ.BmbpRbacOrganUnit{})
	db.DbUtil.AutoMigrate(&organ.BmbpRbacOrganDept{})
	db.DbUtil.AutoMigrate(&organ.BmbpRbacOrganPost{})
	db.DbUtil.AutoMigrate(&organ.BmbpRbacOrganPerson{})
}
