package handler

import (
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/kerokerogeorge/go-gacha-api/internals/usecase"
)

type UserHandler interface {
	Create(*gin.Context)
	GetOne(*gin.Context)
	GetUsers(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
	GetUserCharacters(*gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) *userHandler {
	return &userHandler{
		userUsecase: uu,
	}
}

type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
}

// ユーザ情報作成
func (uh *userHandler) Create(c *gin.Context) {
	var input CreateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name field required"})
		return
	}

	token, err := uh.userUsecase.Create(input.Name)
	if err != nil {
		log.Println(err, gin.H{"error": err})
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// ユーザ情報を一件取得
func (uh *userHandler) GetOne(c *gin.Context) {
	key := c.Request.Header.Get("x-token")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token required"})
		return
	}

	user, err := uh.userUsecase.Get(key)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"name": user.Name})
}

// 全ユーザーの取得
func (uh *userHandler) GetUsers(c *gin.Context) {
	users, err := uh.userUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// ユーザ情報を一件更新
func (uh *userHandler) UpdateUser(c *gin.Context) {
	key := c.Request.Header.Get("x-token")

	var input UpdateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		panic(err)
	}

	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token required"})
		return
	}

	user, err := uh.userUsecase.Get(key)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication failed"})
		return
	}

	updatedUser, err := uh.userUsecase.Update(user, input.Name)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedUser.Name})
}

// ユーザーの削除
func (uh *userHandler) DeleteUser(c *gin.Context) {
	key := c.Request.Header.Get("x-token")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token required"})
		return
	}

	user, err := uh.userUsecase.Get(key)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication failed"})
		return
	}

	err = uh.userUsecase.Delete(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Successfully deleted"})
}

// ユーザ所持キャラクター一覧取得
func (uh *userHandler) GetUserCharacters(c *gin.Context) {
	key := c.Request.Header.Get("x-token")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token required"})
		return
	}

	user, err := uh.userUsecase.Get(key)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "authentication failed"})
		return
	}

	results, err := uh.userUsecase.GetUserCharacters(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	sort.Slice(results, func(i, j int) bool { return results[i].ID < results[j].ID })

	c.JSON(http.StatusOK, gin.H{"characters": results})
}

// func GetEmmitionRate(c *gin.Context) {
// 	var req GetEmmitionRateRequest
// 	var characterEmmitionRateResponse []CharacterEmmitionRateResponse

// 	if err := c.ShouldBindQuery(&req); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	if err := database.DB.Table("gachas").Select("character_emmition_rates.id, character_emmition_rates.character_id, characters.name, character_emmition_rates.emission_rate").
// 		Joins("INNER JOIN character_emmition_rates ON character_emmition_rates.gacha_id = ?", req.GachaID).
// 		Joins("INNER JOIN characters ON character_emmition_rates.character_id = characters.id").
// 		Where("gachas.id = ?", req.GachaID).
// 		Scan(&characterEmmitionRateResponse).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
// 		panic(err)
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": characterEmmitionRateResponse})
// }
