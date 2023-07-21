package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kadycui/go-rank/controllers"
	"github.com/kadycui/go-rank/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
}

func main() {
	r := gin.Default()
	r.POST("/post", controllers.PostCreate)
	r.PUT("/post/:id", controllers.PostsUpdate)

	r.GET("/posts", controllers.PostsIndex)
	r.GET("/post/:id", controllers.PostsShow)

	r.DELETE("/post/:id", controllers.PostsDelete)
	r.Run()
}
