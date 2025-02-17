package dao

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var _db *gorm.DB

func Database(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connWrite, 
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大存活时间

	_db = db

	// 配置主从数据库
	err = _db.Use(dbresolver.Register(
		dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(connWrite)},                  // 写操作源
			Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)}, // 读操作副本
			Policy:   dbresolver.RandomPolicy{},                                // 随机选择副本
		}),
	)
	if err != nil {
		panic("failed to register dbresolver")
	}

	Migration()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}