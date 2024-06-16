package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartAllSpiders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Started all scrappers",
	})
}
func StartSpiderWithId(c *gin.Context) {
	spider := c.Param("spiderId")
	c.JSON(http.StatusOK, gin.H{
		"message": "Started spider:" + " " + spider,
	})
}
func StopAllSpiders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Stopped all Scraping",
	})
}
func StopSpiderWithId(c *gin.Context) {
	spider := c.Param("spiderId")
	c.JSON(http.StatusOK, gin.H{
		"message": "Stopped spider:" + " " + spider,
	})
}
