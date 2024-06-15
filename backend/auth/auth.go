package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MyInfo(c *gin.Context) {
	userInfo, _ := c.Get("currentUser")
	c.JSON(http.StatusOK, gin.H{
		"user": userInfo,
	})
}
