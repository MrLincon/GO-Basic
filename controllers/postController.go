package controllers

import (
	"GO-Basic/initializers"
	"GO-Basic/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePost(c *gin.Context) {

	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	if err := c.BindJSON(&body); err != nil { // Use BindJSON for JSON payload
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind request body"})
		return
	}

	if body.Title == "" || body.Body == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and body are required"})
		return
	}

	post := models.Post{Title: body.Body, Body: body.Title}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Post created!",
		"data": gin.H{
			"id":    post.ID,
			"title": post.Title,
			"body":  post.Body,
		},
	})

}

func FetchAllPost(c *gin.Context) {

	var posts []models.Post

	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var responseData []gin.H
	for _, post := range posts {
		responseData = append(responseData, gin.H{
			"id":    post.ID,
			"title": post.Title,
			"body":  post.Body,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "All posts fetched",
		"data":   responseData,
	})
}

func FetchPostById(c *gin.Context) {

	id := c.Param("id")

	var post models.Post

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Post fetched",
		"data": gin.H{
			"id":    post.ID,
			"title": post.Title,
			"body":  post.Body,
		},
	})
}

func UpdatePost(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	if err := c.BindJSON(&body); err != nil { // Use BindJSON for JSON payload
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind request body"})
		return
	}

	var post models.Post

	initializers.DB.First(&post, id)

	result := initializers.DB.Model(&post).Updates(models.Post{

		Title: body.Body,
		Body:  body.Title,
	})

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Post updated",
		"data": gin.H{
			"id":    post.ID,
			"title": post.Title,
			"body":  post.Body,
		},
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	result := initializers.DB.Delete(&post, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Post deleted",
	})
}
