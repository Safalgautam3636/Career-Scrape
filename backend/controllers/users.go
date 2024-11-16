package controllers

import (
	"backend/helpers"
	"backend/models"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// return message/ error
func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	result := models.DB.Where("username = ? OR email= ?", user.Username, user.Username).First(&user)
	if result.RowsAffected == 1 {
		c.JSON(http.StatusOK, gin.H{
			"message": "User or email already registered please login!",
		})
		return
	}
	user.Password = hashedPassword

	err = models.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, err := helpers.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully!", "token": token, "user": gin.H{
		"username": user.Username,
		"email":    user.Email,
		"isAdmin":  user.IsAdmin,
	}})
}

// LOGIN CONTROLLER
func Login(c *gin.Context) {
	var loginUser models.LoginUser
	fmt.Println(c.Params.Get("password"))
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		// fmt.Print("this one is printed...")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	username := loginUser.Username
	var user models.User
	result := models.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"error": result.Error.Error()})
		return
	}

	checkError := helpers.CompareHash(user.Password, loginUser.Password)
	if checkError != nil {
		c.JSON(http.StatusOK, gin.H{"error": checkError.Error()})
		return
	}
	token, tokenError := helpers.GenerateToken(username)
	if tokenError != nil {
		c.JSON(http.StatusOK, gin.H{"error": tokenError.Error()})
		return
	}

	c.Set("currentUser", user)
	c.JSON(http.StatusOK, gin.H{"message": "Sucessfully logged in!", "token": token, "user": gin.H{
		"email":    user.Email,
		"username": user.Username,
		"isAdmin":  user.IsAdmin,
	}})
}

// GET ALL THE USERS..
func GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := models.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// USER with an ID
func GetUserWithId(c *gin.Context) {
	UserId := string(c.Param("id"))
	_, err := uuid.Parse(UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user models.User
	singleUser := models.DB.First(&user, "id = ?", UserId)
	if singleUser.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": singleUser.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func UpdateUserWithId(c *gin.Context) {
	id := c.Param("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.First(&user, "id = ?", userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	// put the value from the request to the newUserNeedAnUpdate
	var newUserNeedAnUpdate models.User
	if err := c.ShouldBindJSON(&newUserNeedAnUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newUserNeedAnUpdate.Password != "" {
		hashedPassword, _ := helpers.HashPassword(newUserNeedAnUpdate.Password)
		newUserNeedAnUpdate.Password = hashedPassword
	}
	if err := models.DB.Model(&user).Updates(newUserNeedAnUpdate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch the updated user data to return a consistent response
	var updatedUser models.User
	if err := models.DB.First(&updatedUser, "id = ?", userId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send successful response with updated user data
	c.JSON(http.StatusOK, gin.H{
		"user": updatedUser,
	})
}

// Delete all users
func DeleteAllUser(c *gin.Context) {
	err := models.DB.Unscoped().Where("1=1").Delete(&models.User{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "All users deleted",
	})
}
func DeleteUserWithId(c *gin.Context) {
	UserId := string(c.Param("id"))
	_, err := uuid.Parse(UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user models.User
	singleUser := models.DB.Unscoped().Where("id", UserId).Delete(&user)
	if singleUser.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": singleUser.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func VerifyTokens(c *gin.Context) {
	var tokenSchema models.VerifyRequestBody
	err := c.BindJSON(&tokenSchema)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid token!",
			"error":   err.Error(),
		})
		return
	}
	fmt.Println(tokenSchema)
	_, err = helpers.VerifyToken(tokenSchema.Token)
	fmt.Println(err)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"valid": true})
		return
	}
	c.JSON(http.StatusOK, gin.H{"valid": false})
}
