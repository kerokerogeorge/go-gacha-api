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
}

type characterHandler struct {
	characterUsecase usecase.CharacterUsecase
}

func NewCharacterHandler(cu usecase.CharacterUsecase) *characterHandler {
	return &characterHandler{
		characterUsecase: cu,
	}
}

type ResultCharacterResponse struct {
	Id   string `json:"userCharacterId"`
	ID   string `json:"characterId"`
	Name string `json:"name"`
}

type CreateCharacterRequest struct {
	Name string `json:"name"`
}

type Character struct {
	CharacterID  string `json:"characterId"`
	Name         string `json:"name"`
	EmissionRate int    `json:"emissionRate"`
}

// 全キャラクターを取得
func (ch *characterHandler) GetCharacters(c *gin.Context) {
	characters, err := ch.characterUsecase.GetCharacters()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": characters})
}

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
