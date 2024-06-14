package controllers

import (
	"backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	var jobs []models.Job
	if err := models.DB.Find(&jobs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"jobs": jobs,
	})
}

func GetJobById(c *gin.Context) {
	JobId := string(c.Param("id"))
	fmt.Println(JobId)
	_, err := uuid.Parse(JobId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Please provide a valid job id!",
		})
		return
	}
	var job models.Job
	singleJob := models.DB.First(&job, "id = ?", JobId)
	if singleJob.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User does not exists!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "this is a single job",
		"job":     job,
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
