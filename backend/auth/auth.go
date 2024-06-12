package auth

import (
	"backend/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
	"github.com/joho/godotenv"
	"os"
)

type SignupUser struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	IsAdmin  bool    `json:"isAdmin" gorm:"default:false"`
}

type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
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
	sampleSecretKey= []byte(os.Getenv("JWT_KEY"))
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

func VerifyToken(tokenString string) (*jwt.Token, error) {
	godotenv.Load()
	sampleSecretKey= []byte(os.Getenv("JWT_KEY"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { return sampleSecretKey, nil })
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return token, nil
}

func Signup(c *gin.Context) {
	var user SignupUser
	//check for json
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
		ID:       uuid.New().String(),
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
		IsAdmin:user.IsAdmin,
	}
	err = models.DB.Create(&u).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save the user...",
		})
	}
	token, err := GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Token has not been generated...",
		})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "token": token})
}

func Login(c *gin.Context) {
	var loginUser LoginUser
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
		c.JSON(http.StatusOK, gin.H{"data": result.Logger})
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
	c.JSON(http.StatusOK, gin.H{"data": "You logged in!", "token": token})

	//password:=user.Password
	// models.DB.Model().Where("username=?",user.Username,)
	// if user exist and if it does then check the password and if it does generate jwt token
}
func MyInfo(c *gin.Context) {
	userInfo, _ := c.Get("currentUser")
	c.JSON(http.StatusOK, gin.H{
		"user": userInfo,
	})
}