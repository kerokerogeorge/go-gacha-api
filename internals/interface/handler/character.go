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

type GetEmmitionRateRequest struct {
	GachaID string `form:"gachaId"`
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

// ガチャ実行API

// ユーザ所持キャラクター一覧取得
// func GetCharacterList(c *gin.Context) {
// 	var user model.User
// 	var results []ResultCharacterResponse

// 	key := c.Request.Header.Get("x-token")
// 	if key == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Token required"})
// 		return
// 	}

// 	if err := database.DB.Table("users").Where("token = ?", key).First(&user).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed"})
// 		panic(err)
// 	}

// 	if err := database.DB.Table("users").Select("user_characters.id, characters.id, characters.name").
// 		Joins("INNER JOIN user_characters ON user_characters.user_id = ?", user.ID).
// 		Joins("INNER JOIN characters ON user_characters.character_id = characters.id").
// 		Where("users.id = ?", user.ID).
// 		Scan(&results).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
// 		panic(err)
// 	}

// 	sort.Slice(results, func(i, j int) bool { return results[i].Id < results[j].Id })

// 	c.JSON(http.StatusOK, gin.H{"characters": results})
// }

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
