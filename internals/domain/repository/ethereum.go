package repository

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
)

type EthereumRepository interface {
	TransferToken(ctx *gin.Context, from string, to string, contract string, transferAmountOfToken *big.Int) (*types.Transaction, error)
}
