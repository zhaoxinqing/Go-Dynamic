package lib

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB // mysql

// MySQL ...
type DBConfigs struct {
	Source   string `yaml:"source" mapstructure:"source"`
	Replica1 string `yaml:"replica1" mapstructure:"replica1"`
	Replica2 string `yaml:"replica2" mapstructure:"replica2"`
}

// InitDB ...
func InitDB(conf *DBConfigs) (*gorm.DB, error) {

	conn, err := gorm.Open(mysql.Open(conf.Source), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             0,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			}),
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		return nil, err
	}

	// 读写分
	err = conn.Use(dbresolver.
		Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(conf.Source)},
			Replicas: []gorm.Dialector{mysql.Open(conf.Replica1), mysql.Open(conf.Replica2)},
			Policy:   dbresolver.RandomPolicy{}, // sources/replicas 负载均衡策略
		}),
	)
	if err != nil {
		return nil, err
	}

	// 连接池
	sqlDB, err := conn.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(100)  // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(1000) // SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(0) // SetConnMaxLifetime 设置了连接可复用的最大时间。
	return conn, nil
}

// var globalTimezone string
// db.DB.Raw("SELECT @@global.time_zone").Scan(&globalTimezone)
// fmt.Println(globalTimezone)

// var sessionTimezone string
// db.DB.Raw("SELECT  @@session.time_zone").Scan(&sessionTimezone)
// fmt.Println(sessionTimezone)

// var consignment model.Consignment
// db.DB.Raw("SELECT * from consignment order by created_at desc limit 1").Scan(&consignment)
// fmt.Println(consignment.CreatedAt)
// fmt.Println(consignment.UpdatedAt)

// // info := model.Whitelist{
// // 	UserID:        56,
// // 	CollectibleID: 1,
// // 	Status:        0,
// // }
// // db.DB.Create(&info)

// // go func() {
// // 	for {
// // 		logx.Info("测试-------------------------------------")
// // 		time.Sleep(time.Second * 5)
// // 	}
// // }()

// 循环
// go func() {
// 	for {
// 		logx.Info("测试-------------------------------------")
// 		time.Sleep(time.Second * 1)
// 	}
// }()
