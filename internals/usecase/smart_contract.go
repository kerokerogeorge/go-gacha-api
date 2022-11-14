package usecase

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/repository"
)

type ContractUsecase interface {
	GetTransferTokenTransactionPayload(ctx *gin.Context, from string, to string, contract string, amount *big.Int) (*types.Transaction, error)
}

type contractUsecase struct {
	ethereumRepo repository.EthereumRepository
}

func NewContractUsecase(
	er repository.EthereumRepository,
) ContractUsecase {
	return &contractUsecase{
		ethereumRepo: er,
	}
}

func (ctu *contractUsecase) GetTransferTokenTransactionPayload(ctx *gin.Context, from string, to string, contract string, amount *big.Int) (*types.Transaction, error) {
	return ctu.ethereumRepo.TransferToken(ctx, from, to, contract, amount)
}
