package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbUtil *gorm.DB

func init() {
	dsn := "host=localhost user=bmbp password=bmbp dbname=bmbp port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	orm, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		panic("failed to connect database")
	}
	DbUtil = orm
	println("init db")
}
