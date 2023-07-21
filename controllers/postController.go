package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kadycui/go-rank/initializers"
	"github.com/kadycui/go-rank/models"
)

func PostCreate(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.Db.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post

	initializers.Db.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostsShow(c *gin.Context) {

	id := c.Param("id")
	var post models.Post

	initializers.Db.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsUpdate(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	id := c.Param("id")

	var post models.Post
	initializers.Db.First(&post, id)

	initializers.Db.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})

}



func PostsDelete(c *gin.Context) {
	id := c.Param("id")
	initializers.Db.Delete(&models.Post{}, id)


	c.Status(200)

}

