package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	database "github.com/kerokerogeorge/go-test-prod/database"
	"github.com/kerokerogeorge/go-test-prod/model"
)

type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
}

// ユーザ情報作成
func CreateUser(c *gin.Context) {
	var input CreateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err, gin.H{"error": err})
		return
	}

	user := model.User{Name: input.Name, Token: token.String()}
	db := database.DB.Table("users").Create(&user)
	if db.Error != nil {
		panic(db.Error)
	}
	c.JSON(http.StatusOK, gin.H{"token": user.Token})
}

// ユーザ情報取得
func GetUser(c *gin.Context) {
	var user model.User
	key := c.Request.Header.Get("x-token")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token required"})
		return
	}

	if err := database.DB.Table("users").Where("token = ?", key).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed"})
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"name": user.Name})
}

// ユーザ情報更新API
func UpdateUser(c *gin.Context) {
	var user model.User
	key := c.Request.Header.Get("x-token")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token required"})
		return
	}

	if err := database.DB.Table("users").Where("token = ?", key).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed"})
		panic(err)
	}

	var input UpdateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		panic(err)
	}

	db := database.DB.Model(&user).Updates(input)
	if db.Error != nil {
		panic(db.Error)
	}

	c.JSON(http.StatusOK, gin.H{"data": user.Name})
}

// ============
// 以下開発用
// ============

// 全ユーザーの取得
func GetUsers(c *gin.Context) {
	var users []model.User

	if err := database.DB.Find(&users).Error; err != nil {
		panic(err)
	}

	if len(users) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"data": "No user exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// ユーザーの削除
func DeleteUser(c *gin.Context) {
	var user model.User
	key := c.Request.Header.Get("x-token")

	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token required"})
		return
	}

	if err := database.DB.Table("users").Where("token = ?", key).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed"})
		panic(err)
	}

	db := database.DB.Delete(&user)
	if db.Error != nil {
		panic(db.Error)
	}
	c.JSON(http.StatusOK, gin.H{"data": "Successfully deleted"})
}
