package controllers

import (
	"github.com/gaurangkudale/go-crud/initializers"
	"github.com/gaurangkudale/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//get data off request body

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Response with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	//get the id off url
	id := c.Param("id")
	//get the posts
	var posts models.Post
	initializers.DB.First(&posts, id)

	//respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsUpdate(c *gin.Context) {
	//Get the id off the URL
	id := c.Param("id")
	//Get the data of requested body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	//update it
	// Update attributes with `struct`, will only update non-zero fields
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body})
	//respond with it
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsDelete(c *gin.Context) {
	//get the id of the url
	id := c.Param("id")
	//delete the post
	// Email's ID is `10`
	initializers.DB.Delete(&models.Post{}, id)
	// respond
	c.Status(200)
}
