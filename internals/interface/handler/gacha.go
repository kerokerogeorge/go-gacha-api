package handler

import (
	// "math"

	// "math/rand"
	"net/http"
	// "sort"
	// "time"

	// "github.com/Songmu/flextime"
	"github.com/gin-gonic/gin"
	// "github.com/oklog/ulid"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
	"github.com/kerokerogeorge/go-gacha-api/internals/usecase"
)

type GachaHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
}

type gachaHandler struct {
	gachaUsecase usecase.GachaUsecase
}

func NewGachaHandler(gu usecase.GachaUsecase) *gachaHandler {
	return &gachaHandler{
		gachaUsecase: gu,
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

// func GetGacha(c *gin.Context) {
// 	var req GetGachaRequest
// 	var gacha model.Gacha

// 	if err := c.ShouldBindQuery(&req); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := database.DB.Table("gachas").Where("id = ?", req.GachaID).First(&gacha).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
// 		panic(err)
// 	}

// 	characters, err := ToCharacterModel(c, req.GachaID)
// 	if err != nil {
// 		panic(err)
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": characters})
// }

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
