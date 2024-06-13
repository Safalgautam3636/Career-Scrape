package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterJobRoutes(route *gin.Engine) {
	jobs := route.Group("/jobs")
	jobs.GET("/", controllers.GetAllJobs)
	jobs.GET("/:id", controllers.GetJobById)
	jobs.POST("/create",controllers.CreateJob)
	jobs.DELETE("/", controllers.DeleteJobs)
	jobs.DELETE("/:id", controllers.DeleteJobById)
	jobs.PUT("/:id", controllers.UpdateJobById)
}
