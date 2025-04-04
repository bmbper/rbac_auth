package organ

import (
	"api/db"
	"log/slog"
)

func init() {
	slog.Info("初始化组织机构表")
	db.DbUtil.AutoMigrate(&BmbpRbacOrgan{})
	db.DbUtil.AutoMigrate(&BmbpRbacOrganArea{})
	db.DbUtil.AutoMigrate(&BmbpRbacOrganGroup{})
	db.DbUtil.AutoMigrate(&BmbpRbacOrganUnit{})
	db.DbUtil.AutoMigrate(&BmbpRbacOrganDept{})
	db.DbUtil.AutoMigrate(&BmbpRbacOrganPost{})
	db.DbUtil.AutoMigrate(&BmbpRbacOrganPerson{})
}
