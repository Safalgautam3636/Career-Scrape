package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)
// TODO: FUTURE WORK
func RegisterScrapingRoutes(route *gin.Engine) {
	scrape := route.Group("/scrape")
	//start all spiders
	scrape.POST("/start", controllers.StartAllSpiders)
	//start one spider
	scrape.POST("/start/:spiderId", controllers.StartSpiderWithId)
	//stop all spiders
	scrape.POST("/stop", controllers.StopAllSpiders)
	//stop one spider
	scrape.POST("/stop/:spiderId", controllers.StopSpiderWithId)
}
