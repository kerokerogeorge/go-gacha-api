package externals

import (
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/sha3"
)

type ethereumRepository struct {
	ethclient *ethclient.Client
}

func NewEthereumRepository(ethClient *ethclient.Client) *ethereumRepository {
	return &ethereumRepository{
		ethclient: ethClient,
	}
}

func (er *ethereumRepository) TransferToken(ctx *gin.Context, from string, to string, contract string, transferAmountOfToken int) (*types.Transaction, error) {
	fromAddress := common.HexToAddress(from)
	toAddress := common.HexToAddress(to)
	tokenAddress := common.HexToAddress(contract)

	nonce, err := er.ethclient.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, err
	}

	value := big.NewInt(0)

	gasTipCap, err := er.ethclient.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, err
	}

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)

	methodID := hash.Sum(nil)[:4]
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	amount := big.NewInt(int64(transferAmountOfToken))
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	chainID, err := er.ethclient.NetworkID(ctx)
	if err != nil {
		return nil, err
	}

	gasLimit, err := er.ethclient.EstimateGas(ctx, ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
		From: fromAddress,
	})
	if err != nil {
		return nil, err
	}

	block, err := er.ethclient.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	baseFee := block.BaseFee()
	maxFee := baseFee.Mul(baseFee, big.NewInt(2))

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		To:        &tokenAddress,
		Value:     value,
		GasTipCap: gasTipCap,
		GasFeeCap: maxFee,
		Gas:       gasLimit,
		Data:      data,
	})

	return tx, nil
}
