package repository

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
)

type EthereumRepository interface {
	TransferToken(ctx *gin.Context, from string, to string, contract string, transferAmountOfToken *big.Int) (*types.Transaction, error)
	BuyToken(ctx *gin.Context, from string, contract string) (*types.Transaction, error)
	RawTransaction(ctx *gin.Context, from string, to string, contract string, transferAmountOfToken *big.Int) (string, *types.Receipt, error)
	CheckAccountTokenBalance(from string, contract string, transferAmountOfToken *big.Int) (bool, error)
}
