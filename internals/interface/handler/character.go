package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
	"github.com/kerokerogeorge/go-gacha-api/internals/usecase"
)

type CharacterHandler interface {
	GetCharacters(c *gin.Context)
	Create(c *gin.Context)
	GetWithEmmitionRate(c *gin.Context)
	Delete(c *gin.Context)
}

type characterHandler struct {
	characterUsecase usecase.CharacterUsecase
}

func NewCharacterHandler(cu usecase.CharacterUsecase) *characterHandler {
	return &characterHandler{
		characterUsecase: cu,
	}
}

type CreateCharacterRequest struct {
	Name   string `json:"name"`
	ImgUrl string `json:"imgUrl"`
}

// @Summary キャラクター一覧を取得するAPI
// @Router /character/list [get]
// @Description 登録されているキャラクター一覧を取得します
// @Accept application/json
// @Success 200 {object} []model.Character
// @Failure 400 {object} helper.Error
func (ch *characterHandler) GetCharacters(c *gin.Context) {
	characters, err := ch.characterUsecase.List()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": characters})
}

// @Summary キャラクターを作成するAPI
// @Router /character [post]
// @Description 新しいキャラクターを作成します
// @Accept application/json
// @Param name body string true "name"
// @Param imgUrl body string true "imgUrl"
// @Success 200 {object} model.Character
// @Failure 400 {object} helper.Error
func (ch *characterHandler) Create(c *gin.Context) {
	var req CreateCharacterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCharacter, err := model.NewCharacter(req.Name, req.ImgUrl)
	if err != nil {
		panic(err)
	}

	character, err := ch.characterUsecase.Create(newCharacter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": character})
}

// @Summary キャラクター一覧を排出率とともに取得するAPI
// @Router /character/emmition_rates [get]
// @Description キャラクター一覧を排出率とともに取得します
// @Accept application/json
// @Param gachaId path string true "gachaId"
// @Success 200 {object} []model.CharacterWithEmmitionRate
// @Failure 400 {object} helper.Error
func (ch *characterHandler) GetWithEmmitionRate(c *gin.Context) {
	characters, err := ch.characterUsecase.GetCharactersWithEmmitionRate(c.Param("gachaId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": characters})
}

// @Summary キャラクターを削除するAPI
// @Router /character/{characterId} [delete]
// @Description キャラクターを一件削除します
// @Accept application/json
// @Param characterId path string true "characterId"
// @Success 204
// @Failure 400 {object} helper.Error
func (ch *characterHandler) Delete(c *gin.Context) {
	err := ch.characterUsecase.Delete(c.Param("characterId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete character failed"})
		return
	}

	c.Status(http.StatusNoContent)
}
