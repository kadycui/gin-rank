package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

type User struct {
	gorm.Model
	Name string
}

func main() {
	var err error
	dsn := "host=127.0.0.1 user=kadyc password=123456 dbname=test port=5432 TimeZone=Asia/Shanghai"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("db conect fail !!")
	}

	fmt.Println(Db)

	Db.AutoMigrate(&User{})

}
