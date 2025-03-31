package core

import (
	organ "api/core/rbac/organ"
	"api/db"
)

func init() {
	println("init core")
	db.DbUtil.Table("bmbp_rbac_organ").AutoMigrate(&organ.BmbpRbacOrgan{})
}
