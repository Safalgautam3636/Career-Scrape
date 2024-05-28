package main

import (
	// "net/http"
	// "github.com/gin-gonic/gin"
	"backend/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	// router:=gin.Default()
	// router.GET("/",func(c *gin.Context){
	// 	c.JSON(http.StatusOK,gin.H{
	// 		"data":"Hello world"
	// 	})
	// })
	godotenv.Load()
	host:=os.Getenv("DB_HOST")
	username:=os.Getenv("DB_USER")
	password:=os.Getenv("DB_PASSWORD")
	name:=os.Getenv("DB_NAME")
	port:=os.Getenv("DB_PORT")
	// fmt.Println("host",host,username,password,name,port)
	models.ConnectDatabase(host,username,password,name,port)
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