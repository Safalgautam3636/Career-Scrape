package routes

import (
	"fmt"
	"net/http"
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterJobRoutes(route *gin.Engine) {
	jobs := route.Group("/jobs")
	jobs.GET("/", controllers.GetAllJobs)
	jobs.GET("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "this is a single job",
		})
	})
	jobs.POST("/create", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "New job created",
		})
	})
	jobs.DELETE("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "All jobs deleted",
		})
	})
	jobs.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		message := fmt.Sprintf("job with id %s deleted", id)
		c.JSON(http.StatusOK, gin.H{

			"message": message,
		})
	})
	jobs.PUT("/:id", func(c *gin.Context) {
		id := c.Param("id")
		message := fmt.Sprintf("Job with id %s is updated", id)
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	})
}
