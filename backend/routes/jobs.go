package routes

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

// types of user
// normal job seeker and the admin user

func RegisterJobRoutes(route *gin.Engine) {
	jobs := route.Group("/jobs")
	// normal user
	jobs.GET("/", middlewares.CheckAuth, controllers.GetAllJobs)
	// normal user
	jobs.GET("/:id", middlewares.CheckAuth, controllers.GetJobById)
	// super user/admin
	jobs.POST("/create", middlewares.CheckAuth, controllers.CreateJob)
	jobs.DELETE("/", middlewares.CheckAuth, controllers.DeleteJobs)
	jobs.DELETE("/:id", middlewares.CheckAuth, controllers.DeleteJobById)
	jobs.PUT("/:id",middlewares.CheckAuth, controllers.UpdateJobById)
}
