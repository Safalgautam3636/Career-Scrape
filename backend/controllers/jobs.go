package controllers

import (
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateJob(c *gin.Context) {
	var job models.Job
	err := c.ShouldBindJSON(&job)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = models.DB.Create(&job).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "New job created",
	})
}

func GetAllJobs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "All jobs",
	})
}

func GetJobById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "this is a single job",
	})
}

func DeleteJobs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "All jobs deleted",
	})
}

func DeleteJobById(c *gin.Context) {
	id := c.Param("id")
	message := fmt.Sprintf("job with id %s deleted", id)
	c.JSON(http.StatusOK, gin.H{

		"message": message,
	})
}

func UpdateJobById(c *gin.Context) {
	id := c.Param("id")
	message := fmt.Sprintf("Job with id %s is updated", id)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
