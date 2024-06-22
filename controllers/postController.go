package controllers

import (
	"GO-Basic/initializers"
	"GO-Basic/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "title and body are required"})
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
	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	var post models.Post

	result := initializers.DB.First(&post, parsedID)

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
	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	var body struct {
		Body  string
		Title string
	}

	if err := c.BindJSON(&body); err != nil { // Use BindJSON for JSON payload
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind request body"})
		return
	}

	var post models.Post

	initializers.DB.First(&post, parsedID)

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
	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	var post models.Post

	result := initializers.DB.Delete(&post, parsedID)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Post deleted",
	})
}
