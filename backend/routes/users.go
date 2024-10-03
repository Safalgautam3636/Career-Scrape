package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(route *gin.Engine) {
	users := route.Group("/users")
	users.POST("/login", controllers.Login)
	users.POST("/signup", controllers.Signup)
	users.GET("/", controllers.GetAllUsers)
	users.GET("/:id", controllers.GetUserWithId)
	users.DELETE("/", controllers.DeleteAllUser)
	users.DELETE("/:id", controllers.DeleteUserWithId)
	users.PUT("/:id", controllers.UpdateUserWithId)
	users.POST("/verify-token",controllers.VerifyTokens)
}
