package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
	"github.com/kerokerogeorge/go-gacha-api/internals/usecase"
)

type UserCharacterHandler interface {
	Update(*gin.Context)
}

type userCharacterHandler struct {
	userCharacterUsecase usecase.UserCharacterUsecase
}

func NewUserCharacterHandler(ucu usecase.UserCharacterUsecase) *userCharacterHandler {
	return &userCharacterHandler{
		userCharacterUsecase: ucu,
	}
}

type UpdateUserCharacterStatusRequest struct {
	Status           model.CharacterStatus `json:"status"`
	UserCharacterIDs []uint                `json:"userCharacterIds"`
}

// @Summary ユーザー情報を更新するAPI
// @Router /user_character [put]
// @Description ユーザーを一件更新する
// @Accept application/json
// @Param x-token header string true "x-token"
// @Param UserCharacterIDs body []string true "UserCharacterIDs"
// @Success 200 {object} []model.UserCharacter
// @Failure 400 {object} helper.Error
func (uch *userCharacterHandler) Update(c *gin.Context) {
	var req UpdateUserCharacterStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		panic(err)
	}
	key := c.Request.Header.Get("x-token")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token required"})
		return
	}

	if req.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status required"})
		return
	}

	results, err := uch.userCharacterUsecase.UpdateStatus(key, req.Status, req.UserCharacterIDs)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}
