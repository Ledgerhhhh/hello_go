package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	DbL *gorm.DB
)

func SetUpDB() {
	dsn := "root:biaoge666@tcp(localhost:3306)/test_go_connect?parseTime=True"
	var err error
	DbL, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	//数据库连接池对象
	sqlDB, err := DbL.DB()
	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)           // 最小空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接的最大生存时间

}
