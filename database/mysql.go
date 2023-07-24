package database

import (
	"fmt"
	"strings"
	"time"

	"github.com/kadycui/gin-rank/conf"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() {
	dns := strings.Join([]string{conf.MySQL_User, ":", conf.MySQL_Password, "@tcp(", conf.MySQL_Host, ":", conf.MySQL_Port, ")/", conf.MySQL_Db, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")

	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dns,
	}), &gorm.Config{})

	// Error处理
	if err != nil {
		logrus.Fatal("MySQL连接失败!!!", err.Error())
	}

	Db = db

	sqlDB, err := db.DB()
	if err != nil {
		logrus.Fatal("获取连接失败", err.Error())
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("MySQL连接成功!!!")

}

func GetDB() *gorm.DB {
	return Db
}
