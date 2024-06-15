package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "this is your GetAllUsers",
	})
}
func GetUserWithId(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "this is your GetUserWithId",
	})
}
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "this is your login route",
	})
}

func Signup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "this is your Signup",
	})
}
func UpdateUserWithId(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "this is your UpdateUserWithId",
	})
}
func DeleteAllUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "this is your DeleteAllUser",
	})
}
func DeleteUserWithId(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "this is your DeleteUserWithId",
	})
}
