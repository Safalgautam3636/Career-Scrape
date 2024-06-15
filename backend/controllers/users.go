package controllers

import (
	"backend/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var sampleSecretKey []byte

// HELPERS
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GenerateToken(username string) (string, error) {
	godotenv.Load()
	sampleSecretKey = []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(10 * time.Minute).Unix(),
	})

	// SECRET here
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func Signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Password cannot be hashed...",
		})
	}
	u := models.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
		IsAdmin:  user.IsAdmin,
	}
	err = models.DB.Create(&u).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to save the user...",
			"message": err.Error(),
		})
	}
	token, err := GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Token has not been generated...",
		})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "token": token, "user": u})
}

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
