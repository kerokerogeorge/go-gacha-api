package handler

import (
	// "math"

	// "math/rand"
	"log"
	"net/http"

	// "sort"

	"github.com/gin-gonic/gin"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
	"github.com/kerokerogeorge/go-gacha-api/internals/usecase"
)

type GachaHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
}

type GetGachaResponse struct {
	GachaId   string                             `json:"gachaId"`
	Character []*model.CharacterWithEmmitionRate `json:"character"`
}

type GachaListResponse struct {
	ID string `json:"gachaId"`
}

type GetGachaRequest struct {
	GachaId string `form:"gachaId"`
}
type gachaHandler struct {
	gachaUsecase     usecase.GachaUsecase
	characterUsecase usecase.CharacterUsecase
}

func NewGachaHandler(gu usecase.GachaUsecase, cu usecase.CharacterUsecase) *gachaHandler {
	return &gachaHandler{
		gachaUsecase:     gu,
		characterUsecase: cu,
	}
}

func (gh *gachaHandler) Create(c *gin.Context) {
	newGacha, err := model.NewGacha()
	if err != nil {
		panic(err)
	}

	gacha, err := gh.gachaUsecase.Create(newGacha)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": gacha.ID})
}

func (gh *gachaHandler) Get(c *gin.Context) {
	var req GetGachaRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(req.GachaId)
	gacha, err := gh.gachaUsecase.Get(req.GachaId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
		return
	}

	charactersWithEmmitionRate, err := gh.characterUsecase.GetCharactersWithEmmitionRate(gacha.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Characters not found"})
		return
	}

	getGachaResponse := &GetGachaResponse{
		GachaId:   gacha.ID,
		Character: charactersWithEmmitionRate,
	}

	c.JSON(http.StatusOK, gin.H{"data": getGachaResponse})
}

func (gh *gachaHandler) List(c *gin.Context) {
	var res []GachaListResponse
	gachas, err := gh.gachaUsecase.List()
	if err != nil {
		panic(err)
	}

	for _, gacha := range gachas {
		res = append(res, GachaListResponse{ID: gacha.ID})
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

// func DeleteGacha(c *gin.Context) {
// 	var req DeleteGachaRequest
// 	var gacha model.Gacha

// 	if err := c.ShouldBindQuery(&req); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := database.DB.Table("gachas").Where("id = ?", req.GachaID).First(&gacha).Error; err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Record Not Found"})
// 		panic(err)
// 	}

// 	db := database.DB.Delete(&gacha)
// 	if db.Error != nil {
// 		panic(db.Error)
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": "Successfully deleted"})
// }

// func ToCharacterModel(c *gin.Context, gachaId string) (*GetGachaResponse, error) {
// 	var character []*Character
// 	if err := database.DB.Table("gachas").Select("character_emmition_rates.character_id, characters.name, character_emmition_rates.emission_rate").
// 		Joins("INNER JOIN character_emmition_rates ON character_emmition_rates.gacha_id = ?", gachaId).
// 		Joins("INNER JOIN characters ON character_emmition_rates.character_id = characters.id").
// 		Where("gachas.id = ?", gachaId).
// 		Scan(&character).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
// 		panic(err)
// 	}

// 	getGachaResponse := &GetGachaResponse{
// 		GachaID:    gachaId,
// 		Characters: character,
// 	}
// 	return getGachaResponse, nil
// }
