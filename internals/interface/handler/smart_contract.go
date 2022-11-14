package handler

import (
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/kerokerogeorge/go-gacha-api/internals/usecase"
)

type ContractHandler interface {
	GetTransferTokenTransactionPayload(c *gin.Context)
}

type GetTransferTokenTransactionPayloadRequest struct {
	FromAddress     string   `json:"fromAddress"`
	ToAddress       string   `json:"toAddress"`
	ContractAddress string   `json:"contractAddress"`
	Amount          *big.Int `json:"amount"`
}

type TransferTokenTransactionPayloadResponse struct {
	Transaction *types.Transaction `json:"transaction"`
}

type contractHandler struct {
	contractUsecase usecase.ContractUsecase
}

func NewContractHandler(ctu usecase.ContractUsecase) *contractHandler {
	return &contractHandler{
		contractUsecase: ctu,
	}
}

// @Summary ガチャ一覧を取得するAPI
// @Router /contract/transfer [get]
// @Description トークン送金に使用するトランザクションのペイロードを取得します
// @Accept application/json
// @Success 200 {object} []TransferTokenTransactionPayloadResponse
// @Failure 400 {object} helper.Error
func (cth *contractHandler) GetTransferTokenTransactionPayload(c *gin.Context) {
	var req GetTransferTokenTransactionPayloadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payload, err := cth.contractUsecase.GetTransferTokenTransactionPayload(c, req.FromAddress, req.ToAddress, req.ContractAddress, req.Amount)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &TransferTokenTransactionPayloadResponse{
		Transaction: payload,
	})
}
