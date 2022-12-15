package handler

import (
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
	"github.com/kerokerogeorge/go-gacha-api/internals/usecase"
)

type GachaHandler interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Draw(c *gin.Context)
	DrawWithTransaction(c *gin.Context)
	Delete(c *gin.Context)
	ListResult(c *gin.Context)
}

type GetGachaResponse struct {
	GachaId    string                             `json:"gachaId"`
	Characters []*model.CharacterWithEmmitionRate `json:"characters"`
}

type GachaListResponse struct {
	ID string `json:"gachaId"`
}

type CreateGachaResponse struct {
	GachaId string `json:"id"`
}

type DrawGachaRequest struct {
	Times int `json:"times"`
}
type DrawGachaResponse struct {
	Result []*model.Result `json:"result"`
}

type DrawGachaWithTransactionRequest struct {
	Times           int      `json:"times"`
	FromAddress     string   `json:"fromAddress"`
	ToAddress       string   `json:"toAddress"`
	ContractAddress string   `json:"contractAddress"`
	Amount          *big.Int `json:"amount"`
}

type DrawGachaWithTransactionResponse struct {
	Result      []*model.Result `json:"result"`
	Transaction string          `json:"transaction"`
	Receipt     *types.Receipt  `json:"receipt"`
}

type ResultHistoryResponse struct {
	Result []*model.Result `json:"result"`
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
// @Success 200 {object} []GachaListResponse
// @Failure 400 {object} helper.Error
func (gh *gachaHandler) List(c *gin.Context) {
	var res []GachaListResponse
	gachas, err := gh.gachaUsecase.List()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
// @Success 200 {object} CreateGachaResponse
// @Failure 400 {object} helper.Error
func (gh *gachaHandler) Create(c *gin.Context) {
	gacha, err := gh.gachaUsecase.Create()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &CreateGachaResponse{
		GachaId: gacha.ID,
	})
}

// @Summary ガチャを一件取得するAPI
// @Router /gacha/{gachaId} [get]
// @Description 新しいガチャと登録されているキャラクターの排出率を取得する
// @Accept application/json
// @Param gachaId path string true "gachaId"
// @Success 200 {object} model.Gacha
// @Failure 400 {object} helper.Error
func (gh *gachaHandler) Get(c *gin.Context) {
	gacha, err := gh.gachaUsecase.Get(c.Param("gachaId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gacha})
}

// @Summary ガチャを実行するAPI
// @Router /gacha/draw/{gachaId} [post]
// @Description ガチャを実行し、キャラクターを取得します
// @Accept application/json
// @Param x-token header string true "x-token"
// @Param gachaId path string true "gachaId"
// @Param times body string true "ガチャを実行する回数"
// @Success 200 {object} DrawGachaResponse
// @Failure 400 {object} helper.Error
func (gh *gachaHandler) Draw(c *gin.Context) {
	var req DrawGachaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key := c.Request.Header.Get("x-token")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token required"})
		return
	}

	results, err := gh.gachaUsecase.Draw(c, c.Param("gachaId"), req.Times, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, &DrawGachaResponse{
		Result: results,
	})
}

// @Summary ガチャを削除するAPI
// @Router /gacha/{gachaId} [delete]
// @Description ガチャを一件削除します
// @Accept application/json
// @Param gachaId path string true "gachaId"
// @Success 204
// @Failure 400 {object} helper.Error
func (gh *gachaHandler) Delete(c *gin.Context) {
	err := gh.gachaUsecase.Delete(c.Param("gachaId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete failed"})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary キャラ排出履歴一覧取得API
// @Router /gacha/result [get]
// @Description キャラ排出履歴一覧を取得します
// @Accept application/json
// @Success 200 {object} model.Result
// @Failure 400 {object} helper.Error
func (gh *gachaHandler) ListResult(c *gin.Context) {
	key := c.Request.Header.Get("x-token")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token required"})
		return
	}

	results, err := gh.gachaUsecase.ListHistory(c, key)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, &ResultHistoryResponse{
		Result: results,
	})
}

// @Summary ガチャを実行するAPI
// @Router /gacha/draw_with_transaction/{gachaId} [post]
// @Description ガチャを実行し、キャラクターを取得します
// @Accept application/json
// @Param x-token header string true "x-token"
// @Param gachaId path string true "gachaId"
// @Param times body string true "ガチャを実行する回数"
// @Success 200 {object} DrawGachaResponse
// @Failure 400 {object} helper.Error
func (gh *gachaHandler) DrawWithTransaction(c *gin.Context) {
	var req DrawGachaWithTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key := c.Request.Header.Get("x-token")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token required"})
		return
	}

	results, tx, receipt, err := gh.gachaUsecase.DrawWithTransaction(c, c.Param("gachaId"), req.Times, key, req.FromAddress, req.ToAddress, req.ContractAddress, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, &DrawGachaWithTransactionResponse{
		Result:      results,
		Transaction: tx,
		Receipt:     receipt,
	})
}
