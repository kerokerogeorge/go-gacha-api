package repository

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
)

type EthereumRepository interface {
	TransferToken(ctx *gin.Context, from string, to string, contract string, transferAmountOfToken int) (*types.Transaction, error)
}
