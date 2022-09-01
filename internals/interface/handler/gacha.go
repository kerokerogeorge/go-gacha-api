package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
	"github.com/kerokerogeorge/go-gacha-api/internals/usecase"
)

type GachaHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Draw(c *gin.Context)
	Delete(c *gin.Context)
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

type CreateGachaRequest struct {
	Times   int    `json:"times"`
	GachaID string `json:"gachaId"`
}

type GachaResultResponse struct {
	ID   string `json:"characterId"`
	Name string `json:"name"`
}

type DeleteGachaRequest struct {
	GachaId string `form:"gachaId"`
}

type gachaHandler struct {
	gachaUsecase     usecase.GachaUsecase
	characterUsecase usecase.CharacterUsecase
	userUsecase      usecase.UserUsecase
}

func NewGachaHandler(gu usecase.GachaUsecase, cu usecase.CharacterUsecase, uu usecase.UserUsecase) *gachaHandler {
	return &gachaHandler{
		gachaUsecase:     gu,
		characterUsecase: cu,
		userUsecase:      uu,
	}
}

// @Summary ガチャ一覧を取得するAPI
// @Router /gacha/list [get]
// @Description ガチャ一覧を取得します
// @Accept application/json
// @Success 200 {object} string
// @Failure 400 {object} helper.Error
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

// @Summary 新しいガチャを作成するAPI
// @Router /gacha [post]
// @Description 新しいガチャを作成し、排出率をキャラクターに割り当てます
// @Accept application/json
// @Success 200 {object} string
// @Failure 400 {object} helper.Error
func (gh *gachaHandler) Create(c *gin.Context) {
	gacha, err := gh.gachaUsecase.Create()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"gachaId": gacha.ID})
}

// @Summary ガチャを一件取得するAPI
// @Router /gacha [get]
// @Description 新しいガチャと登録されているキャラクターの排出率を取得する
// @Accept application/json
// @Param gachaId query string true "gachaId"
// @Success 200 {object} GetGachaResponse
// @Failure 400 {object} helper.Error
func (gh *gachaHandler) Get(c *gin.Context) {
	var req GetGachaRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

// @Summary ガチャを実行するAPI
// @Router /gacha/draw [post]
// @Description ガチャを実行し、キャラクターを取得します
// @Accept application/json
// @Param x-token header string true "x-token"
// @Param gachaId body string true "gachaId"
// @Param times body string true "ガチャを実行する回数"
// @Success 200 {object} []GachaResultResponse
// @Failure 400 {object} helper.Error
func (gh *gachaHandler) Draw(c *gin.Context) {
	var req CreateGachaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key := c.Request.Header.Get("x-token")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token required"})
		return
	}

	user, err := gh.userUsecase.Get(key)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	charactersWithEmmitionRate, err := gh.characterUsecase.GetCharactersWithEmmitionRate(req.GachaID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Characters not found"})
		return
	}

	var results []*GachaResultResponse
	for i := 0; i < req.Times; i++ {
		character, err := gh.gachaUsecase.Draw(charactersWithEmmitionRate, user.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "draw gacha failed"})
			return
		}
		// numと配列に格納したN番目の数字をnumに足した値の範囲にランダムに取得した値が含まれていれば、キャラクターIDをもとにキャラクターをDBから取得
		res := &GachaResultResponse{ID: character.ID, Name: character.Name}
		results = append(results, res)
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}

// @Summary ガチャを削除するAPI
// @Router /gacha [delete]
// @Description ガチャを一件削除します
// @Accept application/json
// @Param gachaId query string true "gachaId"
// @Success 204
// @Failure 400 {object} helper.Error
func (gh *gachaHandler) Delete(c *gin.Context) {
	var req DeleteGachaRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gachaCharacters, err := gh.gachaUsecase.GetGachaCharacters(req.GachaId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record Not Found"})
		return
	}

	err = gh.gachaUsecase.DeleteGachaCharacters(gachaCharacters)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete gacha characters failed"})
		return
	}

	gacha, err := gh.gachaUsecase.Get(req.GachaId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gacha record Not Found"})
		return
	}

	err = gh.gachaUsecase.Delete(gacha)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Successfully deleted"})
}
