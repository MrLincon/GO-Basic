package routes

import (
	"GO-Basic/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to GO-Basic",
		})
	})

	postRoutes(r)

}

func postRoutes(r *gin.Engine) {

	postGroup := r.Group("/posts")
	{
		postGroup.POST("/create-post", controllers.CreatePost)
		postGroup.PATCH("/update-post/:id", controllers.UpdatePost)
		postGroup.GET("/fetch-all-post", controllers.FetchAllPost)
		postGroup.GET("/fetch-post/:id", controllers.FetchPostById)
		postGroup.DELETE("/delete-post/:id", controllers.DeletePost)

	}
}
