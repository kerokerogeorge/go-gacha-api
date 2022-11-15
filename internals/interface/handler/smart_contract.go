package handler

import (
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/kerokerogeorge/go-gacha-api/internals/usecase"
)

type ContractHandler interface {
	TransferTokenTransactionPayload(c *gin.Context)
	BuyTransactionPayload(c *gin.Context)
}

type GetTransferTokenTransactionPayloadRequest struct {
	FromAddress     string   `json:"fromAddress"`
	ToAddress       string   `json:"toAddress"`
	ContractAddress string   `json:"contractAddress"`
	Amount          *big.Int `json:"amount"`
}

type GetBuyTokenTransactionPayloadRequest struct {
	FromAddress     string `json:"fromAddress"`
	ContractAddress string `json:"contractAddress"`
}

type TransferTokenTransactionPayloadResponse struct {
	TransactionPayload *types.Transaction `json:"transactionPayload"`
}

type contractHandler struct {
	contractUsecase usecase.ContractUsecase
}

func NewContractHandler(ctu usecase.ContractUsecase) *contractHandler {
	return &contractHandler{
		contractUsecase: ctu,
	}
}

// @Summary トークン送金に使用するトランザクションのペイロードを取得するAPI
// @Router /contract/transfer [post]
// @Description トークン送金に使用するトランザクションのペイロードを取得します
// @Accept application/json
// @Success 200 {object} []TransferTokenTransactionPayloadResponse
// @Failure 400 {object} helper.Error
func (cth *contractHandler) TransferTokenTransactionPayload(c *gin.Context) {
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
		TransactionPayload: payload,
	})
}

// @Summary トークン購入に使用するトランザクションのペイロードを取得するAPI
// @Router /contract/buy [post]
// @Description トークン購入に使用するトランザクションのペイロードを取得します
// @Accept application/json
// @Success 200 {object} []TransferTokenTransactionPayloadResponse
// @Failure 400 {object} helper.Error
func (cth *contractHandler) BuyTransactionPayload(c *gin.Context) {
	var req GetBuyTokenTransactionPayloadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payload, err := cth.contractUsecase.GetBuyTokenTransactionPayload(c, req.FromAddress, req.ContractAddress)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &TransferTokenTransactionPayloadResponse{
		TransactionPayload: payload,
	})
}
