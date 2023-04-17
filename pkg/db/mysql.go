package db

import (
	"backend-go/config"
	"backend-go/public"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB // mysql

// InitDB ...
func InitDB() {
	conf := config.EnvInfo().MySQL

	tempDb, err := gorm.Open(mysql.Open(conf.Source), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             0,
				LogLevel:                  logger.Warn,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			}),
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		panic(fmt.Sprintf("%s - MySQL 连接失败", public.FormatTime()))
	}

	// 读写分
	err = tempDb.Use(dbresolver.
		Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(conf.Source)},
			Replicas: []gorm.Dialector{mysql.Open(conf.Replica1), mysql.Open(conf.Replica2)},
			Policy:   dbresolver.RandomPolicy{}, // sources/replicas 负载均衡策略
		}),
	)
	if err != nil {
		panic(fmt.Sprintf("%s - MySQL 连接失败", public.FormatTime()))
	}

	// 连接池
	sqlDB, err := tempDb.DB()
	if err != nil {
		panic(fmt.Sprintf("%s - 数据库连接池初始化失败!!!", public.FormatTime()))
	}
	sqlDB.SetMaxIdleConns(100)  // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)  // SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(0) // SetConnMaxLifetime 设置了连接可复用的最大时间，0永久

	DB = tempDb // 全局连接

	fmt.Println("MySQL 已连接 !!!")
}
