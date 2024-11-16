package main

import (
	// "backend/authjwt"
	// "backend/auth"
	// "backend/middlewares"
	"backend/middlewares"
	"backend/models"
	"backend/routes"
	"fmt"
	// "backend/auth"
	//"net/http"
	//"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// load env
	godotenv.Load()
	if err := godotenv.Load(); err != nil {
     log.Fatalf("Error loading .env file: %v", err)
	}


	// setup the db
	// "host=plexiform-muse-438603-j7:us-central1:careerscrape user=postgres password=yournewpass dbname=CareerScrape sslmode=disable"
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	models.ConnectDatabase(host,username,password,name,db_port)

	//config routes
	router := gin.Default()
	router.Use(loggingMiddleware(),middlewares.CORSMiddleware())

	routes.RegisterJobRoutes(router)
	routes.RegisterUserRoutes(router)
	// routes.RegisterScrapingRoutes(router)

	serverIp:="0.0.0.0" // "localhost"
	port:="8000"
	log.Printf("API Gateway is running on port %s", port)
	if err := router.Run(fmt.Sprintf("%s:%s",serverIp, port)); err != nil {
		log.Fatalf("Failed to start API Gateway: %s", err.Error())
	}

	router.Run(fmt.Sprintf("%s:%s", "localhost", port))
}
func loggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Request: %s %s", c.Request.Method, c.Request.URL)
		c.Next()
	}
}
