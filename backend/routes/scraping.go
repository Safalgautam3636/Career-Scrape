package routes

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func RegisterScrapingRoutes(route *gin.Engine) {
	scrape := route.Group("/scrape")
	scrape.POST("/start", func(c *gin.Context) {
		go startScraping()
		c.JSON(http.StatusOK, gin.H{
			"message": "Started all scrappers",
		})

	})
	scrape.POST("/:spider", func(c *gin.Context) {
		spider := c.Param("spider")
		go startSingleSpider(spider)
		c.JSON(http.StatusOK, gin.H{
			"message": "Started spider:" + " " + spider,
		})
	})
}
func startScraping() {
	fmt.Println("Started all scraping...")
}
func startSingleSpider(spider string) {
	message := fmt.Sprintf("Spider with name: %s started..", spider)
	fmt.Println(message)
}
