package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllJobs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "All jobs",
	})
}
