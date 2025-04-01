package organ

import (
	"api/db"
	"log/slog"
)

func init() {
	slog.Info("初始化组织机构表")
	db.DbUtil.Table("bmbp_rbac_organ").AutoMigrate(&Organ{})
	db.DbUtil.Table("bmbp_rbac_organ_area").AutoMigrate(&OrganArea{})
	db.DbUtil.Table("bmbp_rbac_organ_group").AutoMigrate(&OrganGroup{})
	db.DbUtil.Table("bmbp_rbac_organ_unit").AutoMigrate(&OrganUnit{})
	db.DbUtil.Table("bmbp_rbac_organ_department").AutoMigrate(&OrganDepartment{})
	db.DbUtil.Table("bmbp_rbac_organ_position").AutoMigrate(&OrganPosition{})
	db.DbUtil.Table("bmbp_rbac_organ_person").AutoMigrate(&OrganPerson{})
}
