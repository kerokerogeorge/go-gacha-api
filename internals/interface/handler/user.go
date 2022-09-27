package handler

import (
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
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
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

type GetUserResponse struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type UpdateUserResponse struct {
	Name string `json:"name"`
}

type UserListResponse struct {
	Users []*model.User `json:"users"`
}

// @Summary ユーザー一覧を取得するAPI
// @Router /user/list [get]
// @Description ユーザー一覧を取得します
// @Accept application/json
// @Success 200 {object} UserListResponse
// @Failure 400 {object} helper.Error
func (uh *userHandler) GetUsers(c *gin.Context) {
	users, err := uh.userUsecase.List()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &UserListResponse{
		Users: users,
	})
}

// @Summary 新しいユーザーを作成するAPI
// @Router /user [post]
// @Description 新しいユーザーを作成します
// @Accept application/json
// @Param name body string true "name"
// @Param address body string true "address"
// @Success 201 {object} CreateUserResponse
// @Failure 400 {object} helper.Error
func (uh *userHandler) Create(c *gin.Context) {
	var input CreateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name field required"})
		return
	}

	token, err := uh.userUsecase.Create(input.Name, input.Address)
	if err != nil {
		log.Println(err, gin.H{"error": err})
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, &CreateUserResponse{
		Token: token,
	})
}

// @Summary 新しいユーザーを一件取得するAPI
// @Router /user [get]
// @Description ユーザーを一件取得する
// @Accept application/json
// @Param x-token header string true "x-token"
// @Success 200 {object} GetUserResponse
// @Failure 400 {object} helper.Error
func (uh *userHandler) GetOne(c *gin.Context) {
	key := c.Request.Header.Get("x-token")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token required"})
		return
	}

	user, err := uh.userUsecase.Get(key)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, &GetUserResponse{
		Name:    user.Name,
		Address: user.Address,
	})
}

// @Summary ユーザー情報を更新するAPI
// @Router /user [put]
// @Description ユーザーを一件更新する
// @Accept application/json
// @Param x-token header string true "x-token"
// @Param name body string true "ユーザー名"
// @Success 200 {object} UpdateUserResponse
// @Failure 400 {object} helper.Error
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

	c.JSON(http.StatusOK, &UpdateUserResponse{
		Name: updatedUser.Name,
	})
}

// @Summary ユーザー情報を削除するAPI
// @Router /user [delete]
// @Description ユーザーを一件削除する
// @Accept application/json
// @Param x-token header string true "x-token"
// @Success 204
// @Failure 400 {object} helper.Error
func (uh *userHandler) DeleteUser(c *gin.Context) {
	key := c.Request.Header.Get("x-token")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token required"})
		return
	}

	err := uh.userUsecase.Delete(key)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary ユーザー所持キャラクター一覧を取得するAPI
// @Router /user/characters [get]
// @Description ユーザー所持キャラクター一覧を取得します
// @Accept application/json
// @Param x-token header string true "x-token"
// @Success 200 {object} []model.Result
// @Failure 400 {object} helper.Error
func (uh *userHandler) GetUserCharacters(c *gin.Context) {
	key := c.Request.Header.Get("x-token")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token required"})
		return
	}

	results, err := uh.userUsecase.GetUserCharacters(key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	sort.Slice(results, func(i, j int) bool { return results[i].ID < results[j].ID })

	c.JSON(http.StatusOK, gin.H{"characters": results})
}
