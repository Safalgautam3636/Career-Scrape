package middlewares

import (
	"backend/helpers"
	"backend/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// TODO: CHECK for token expiry
func CheckAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Auth header is missing..",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token format",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	tokenString := authToken[1]
	token, err := helpers.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Provided token is invalid..",
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}
	username := claims["username"]
	var user models.User
	err = models.DB.Where("username= ?", username).First(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized user..."})
		c.Abort()
	}
	c.Set("currentUser", user)
	c.Next()
}

// func CheckAdmin(c *gin.Context) {
// 	currentUser, _ := c.Get("currentUser")
// 	if !currentUser.isAdmin {
// 		c.JSON(http.StatusPermanentRedirect, gin.H{
// 			"message": "Sorry you dont have enough permissions!",
// 		})
// 		return
// 	}
// 	c.Next()
// }
