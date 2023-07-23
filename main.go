package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kadycui/gin-rank/api"
	"github.com/kadycui/gin-rank/database"
	_ "github.com/kadycui/gin-rank/docs"
	"github.com/kadycui/gin-rank/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			GinRank API
// @version		1.0
// @description	这是一个使用Gin框架实现的系统.
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8080
// @BasePath		/
func main() {

	// 连接数据库
	database.InitDB()

	// 连接redis
	database.InitRedis()

	// 数据迁移
	database.Migration()

	r := gin.Default()

	// 跨域
	r.Use(middleware.Cors())

	// 日志
	r.Use(middleware.LoggerToFile())

	v1 := r.Group("/api/v1")
	{
		v1.POST("auth/register", api.Register)
		v1.POST("auth/login", api.Login)
		v1.GET("auth/info", middleware.AuthMiddleware(), api.Info)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":8080")
	if err != nil {
		return
	}

}
