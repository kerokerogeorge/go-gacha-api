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
	Name string `json:"name"`
}

type GetCharactersWithEmmitionRateRequest struct {
	GachaID string `form:"gachaId"`
}

// @Summary キャラクター一覧を取得するAPI
// @Router /character/list [get]
// @Description 登録されているキャラクター一覧を取得します
// @Accept application/json
// @Success 200 {object} []model.Character
// @Failure 400 {object} helper.Error
func (ch *characterHandler) GetCharacters(c *gin.Context) {
	characters, err := ch.characterUsecase.GetCharacters()
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
// @Success 200 {object} model.Character
// @Failure 400 {object} helper.Error
func (ch *characterHandler) Create(c *gin.Context) {
	var req CreateCharacterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newCharacter, err := model.NewCharacter(req.Name)
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
// @Param gachaId query string true "gachaId"
// @Success 200 {object} []model.CharacterWithEmmitionRate
// @Failure 400 {object} helper.Error
func (ch *characterHandler) GetWithEmmitionRate(c *gin.Context) {
	var req GetCharactersWithEmmitionRateRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	characters, err := ch.characterUsecase.GetCharactersWithEmmitionRate(req.GachaID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": characters})
}

// // もう使わないAPI
// // キャラクターの排出率の変更
// func UpdateCharacter(c *gin.Context) {
// 	var character model.Character

// 	if err := database.DB.Table("characters").Where("id = ?", c.Param("id")).First(&character).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed"})
// 		return
// 	}

// 	var input UpdateCharacterRequest
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	db := database.DB.Model(&character).Updates(input)
// 	if db.Error != nil {
// 		panic(db.Error)
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": character})
// }
