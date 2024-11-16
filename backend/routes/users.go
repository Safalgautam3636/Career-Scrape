package routes

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(route *gin.Engine) {
	users := route.Group("/users")
	//normal user
	users.POST("/login", controllers.Login)
	users.POST("/signup", controllers.Signup)
	//admin
	users.GET("/", middlewares.CheckAuth, controllers.GetAllUsers)
	users.GET("/:id", middlewares.CheckAuth, controllers.GetUserWithId)
	// admin
	users.DELETE("/", middlewares.CheckAuth, controllers.DeleteAllUser)
	users.DELETE("/:id", middlewares.CheckAuth, controllers.DeleteUserWithId)
	// admin or itself
	users.PUT("/:id", middlewares.CheckAuth, controllers.UpdateUserWithId)
	//norml user
	users.POST("/verify-token", middlewares.CheckAuth, controllers.VerifyTokens)
}
