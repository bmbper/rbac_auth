package db

import (
	"bmbp/base"
	"fmt"
	"log/slog"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DbUtil *gorm.DB

func InitDB() error {
	slog.Info("初始化数据库连接...")

	// 使用GlobalConfig中的数据库配置
	dbConfig := base.GlobalConfig.Database
	dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + fmt.Sprintf("%d", dbConfig.Port) + ")/" + dbConfig.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"

	orm, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if dbErr != nil {
		slog.Error("数据库连接失败", "error", dbErr)
		return dbErr
	}

	sqlDB, err := orm.DB()
	if err != nil {
		slog.Error("获取数据库实例失败", "error", err)
		return err
	}

	// 设置连接池
	sqlDB.SetMaxIdleConns(dbConfig.PoolSize / 2)
	sqlDB.SetMaxOpenConns(dbConfig.PoolSize)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DbUtil = orm
	slog.Info("数据库连接成功....")
	return nil
}
