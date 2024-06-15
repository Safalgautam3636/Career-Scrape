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

func Login(c *gin.Context) {
	var loginUser models.LoginUser
	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	username := loginUser.Username
	var user models.User
	result := models.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"data": result.Logger, "message": "Not a valid username!"})
		return
	}

	checkError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if checkError != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Enter Valid password!"})
		return
	}
	token, err := GenerateToken(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"data": "No token hasnot been generated.."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "You logged in!", "token": token, "user": user})

	//password:=user.Password
	// models.DB.Model().Where("username=?",user.Username,)
	// if user exist and if it does then check the password and if it does generate jwt token
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
