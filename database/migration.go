package database

import "github.com/kadycui/gin-rank/model"

// 数据迁移

func Migration() {
	Db.AutoMigrate(&model.User{})
}
