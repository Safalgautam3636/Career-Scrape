package controllers

import (
	"backend/helpers"
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func Signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Password cannot be hashed...",
		})
	}
	result := models.DB.Where("username = ? OR email= ?", user.Username, user.Username).First(&user)
	if result.RowsAffected == 1 {
		c.JSON(http.StatusOK, gin.H{
			"message": "User or email already registered please login!",
		})
		return
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
	token, err := helpers.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Token has not been generated...",
		})
	}
	returnUser := models.ReturnUser{
		Username: u.Username,
		Email:    u.Email,
		IsAdmin:  u.IsAdmin,
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "token": token, "user": returnUser})
}

func Login(c *gin.Context) {
	var loginUser models.LoginUser
	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Server Error!",
		})
		return
	}
	username := loginUser.Username
	var user models.User
	result := models.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Not a valid username!"})
		return
	}

	checkError := helpers.CompareHash(user.Password, loginUser.Password)
	if checkError != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Not a valid password!"})
		return
	}
	token, err := helpers.GenerateToken(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "No token hasnot been generated.."})
		return
	}
	var returnUser = models.ReturnUser{
		Email:    user.Email,
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
	}

	c.Set("currentUser", user)
	c.JSON(http.StatusOK, gin.H{"message": "Sucessfully logged in!", "token": token, "user": returnUser})

	//password:=user.Password
	// models.DB.Model().Where("username=?",user.Username,)
	// if user exist and if it does then check the password and if it does generate jwt token
}

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
func GetUserWithId(c *gin.Context) {
	UserId := string(c.Param("id"))
	_, err := uuid.Parse(UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Please provide a valid user id!",
		})
		return
	}
	var user models.User
	singleUser := models.DB.First(&user, "id = ?", UserId)
	if singleUser.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User does not exists!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "this is a user with the given id",
		"user":    user,
	})
}
func UpdateUserWithId(c *gin.Context) {
	id := c.Param("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}

	var user models.User
	if err := models.DB.First(&user, "id = ?", userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	var newUserNeedAnUpdate models.User
	if err := c.ShouldBindJSON(&newUserNeedAnUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	checkError := helpers.CompareHash(user.Password, newUserNeedAnUpdate.Password)
	fmt.Println(checkError)
	hashedPassword := user.Password
	if checkError != nil {
		hashedPassword, _ = helpers.HashPassword(newUserNeedAnUpdate.Password)

	}
	newUserNeedAnUpdate.Password = hashedPassword
	if err := models.DB.Model(&user).Select("is_admin", "password", "email", "username").Updates(newUserNeedAnUpdate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update job"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Everything is going good!",
		"newUser": newUserNeedAnUpdate,
	})
}
func DeleteAllUser(c *gin.Context) {
	err := models.DB.Where("1 = 1").Delete(&models.User{}).Error
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
			"message": "Please provide a valid user id!",
		})
		return
	}
	var user models.User
	singleUser := models.DB.Where("id", UserId).Delete(&user)
	if singleUser.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User does not exists!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
		"user":    user,
	})
}

func VerifyTokens(c *gin.Context) {
	var tokenSchema models.VerifyRequestBody
	err := c.BindJSON(&tokenSchema)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid token!",
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
