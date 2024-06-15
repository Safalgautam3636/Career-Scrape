package controllers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// func (j *models.Job) AfterDelete(tx *gorm.DB) (err error) {
//   if j.Link=="http://linkedin.com/google-apply" {
//     tx.Model(&models.Job{}).Where("id = ?", j.ID).Delete("invalid", false)
//   }
//   return
// }

func DeleteJobs(c *gin.Context) {
	err := models.DB.Where("1 = 1").Delete(&models.Job{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "All jobs deleted",
	})

}

func DeleteJobById(c *gin.Context) {
	JobId := string(c.Param("id"))
	_, err := uuid.Parse(JobId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Please provide a valid job id!",
		})
		return
	}
	var job models.Job
	singleJob := models.DB.Where("id", JobId).Delete(&job)
	if singleJob.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Job does not exists!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Job deleted",
		"job":     job,
	})
}

func UpdateJobById(c *gin.Context) {
	id := c.Param("id")
	jobID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid job ID"})
		return
	}

	var job models.Job
	if err := models.DB.First(&job, "id = ?", jobID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Job not found"})
		return
	}
	var newJobNeedAnUpdate models.Job
	if err := c.ShouldBindJSON(&newJobNeedAnUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	if err := models.DB.Model(&job).Updates(newJobNeedAnUpdate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update job"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Job updated successfully", "job": job})
}
