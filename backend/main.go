package main

import (
	//"backend/authjwt"
	// "backend/auth"
	// "backend/middlewares"
	"backend/middlewares"
	"backend/models"
	"backend/routes"

	//"backend/auth"

	//"net/http"
	//"net/http"

	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	// load env
	godotenv.Load()

	//setup the db
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	models.ConnectDatabase(host, username, password, name, port)

	//config routes

	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())

	//api:=router.Group("/api")
	// api.POST("/signup",auth.Signup)
	// api.POST("/login",auth.Login)
	// api.GET("/me",middlewares.CheckAuth,auth.MyInfo)
	// {
	// 	api.POST("/register",authjwt.signup)
	// }

	routes.RegisterJobRoutes(router)
	routes.RegisterUserRoutes(router)
	//routes.RegisterScrapingRoutes(router)

	router.Run("localhost:8000")

	// job := models.Job{
	//     Role: "Software Engineer",
	//     Benefits: "Develop software applications",
	// }

	// // Save the job record to the database
	// result := models.DB.Create(&job)
	// if result.Error != nil {
	//     panic("Failed to create job record!")
	// }

	// fmt.Println("Job record created successfully!")

}
