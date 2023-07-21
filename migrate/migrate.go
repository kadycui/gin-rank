package main

import (
	"fmt"

	"github.com/kadycui/go-rank/initializers"
	"github.com/kadycui/go-rank/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
}

func main() {
	fmt.Println("开始迁移数据库")
	initializers.Db.AutoMigrate(&models.Post{})

	fmt.Println("移数完成")
}
