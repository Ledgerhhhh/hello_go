package conn

import (
	"com.ledger.goproject/myconfig"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var DB *sqlx.DB

func InitMysqlConfig() error {
	// 定义数据库连接信息
	dsn := myconfig.GConfig.MysqlConfig.Dsn
	// 创建数据库连接
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("mysql error: %s\n", err)
	}
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("mysql error: %s\n", err)
	}

	db.SetMaxIdleConns(100)          // 配置最大连接数
	db.SetMaxOpenConns(100)          // 配置最大打开的连接数
	db.SetConnMaxLifetime(time.Hour) // 配置连接的最大生命周期
	DB = db
	return nil
}
