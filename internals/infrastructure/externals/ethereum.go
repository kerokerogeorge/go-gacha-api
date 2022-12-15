package externals

import (
	"errors"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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

func (er *ethereumRepository) TransferToken(ctx *gin.Context, from string, to string, contract string, transferAmountOfToken *big.Int) (*types.Transaction, error) {
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

	amount := transferAmountOfToken
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

func (er *ethereumRepository) BuyToken(ctx *gin.Context, from string, contract string) (*types.Transaction, error) {
	fromAddress := common.HexToAddress(from)
	tokenAddress := common.HexToAddress(contract)

	nonce, err := er.ethclient.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, err
	}

	value := big.NewInt(10000000000000)

	gasTipCap, err := er.ethclient.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, err
	}

	signature := []byte("buyTokens()")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(signature)

	methodID := hash.Sum(nil)[:4]

	var data []byte
	data = append(data, methodID...)

	chainID, err := er.ethclient.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	gasLimit, err := er.ethclient.EstimateGas(ctx, ethereum.CallMsg{
		To:    &tokenAddress,
		Data:  data,
		From:  fromAddress,
		Value: value,
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

func (er *ethereumRepository) RawTransaction(ctx *gin.Context, from string, to string, contract string, transferAmountOfToken *big.Int) (string, *types.Receipt, error) {
	log.Println(transferAmountOfToken)
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return "", nil, err
	}
	fromAddress := common.HexToAddress(from)
	toAddress := common.HexToAddress(to)
	tokenAddress := common.HexToAddress(contract)

	nonce, err := er.ethclient.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", nil, err
	}

	value := big.NewInt(0)

	gasTipCap, err := er.ethclient.SuggestGasTipCap(ctx)
	if err != nil {
		return "", nil, err
	}

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)

	methodID := hash.Sum(nil)[:4]
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	amount := transferAmountOfToken
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	chainID, err := er.ethclient.NetworkID(ctx)
	if err != nil {
		return "", nil, err
	}

	gasLimit, err := er.ethclient.EstimateGas(ctx, ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
		From: fromAddress,
	})
	if err != nil {
		return "", nil, err
	}

	block, err := er.ethclient.BlockByNumber(ctx, nil)
	if err != nil {
		return "", nil, err
	}

	baseFee := block.BaseFee()
	estimatedMaxFee := baseFee.Mul(baseFee, big.NewInt(2))

	maxGasFeeUint64 := estimatedMaxFee.Uint64()
	maxPriorityFeeUint64 := gasTipCap.Uint64()

	var maxFeePerGas *big.Int
	if maxGasFeeUint64 > maxPriorityFeeUint64 {
		maxFeePerGas = estimatedMaxFee
	} else {
		maxFeePerGas = gasTipCap
	}
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		To:        &tokenAddress,
		Value:     value,
		GasTipCap: gasTipCap,
		GasFeeCap: maxFeePerGas,
		Gas:       gasLimit,
		Data:      data,
	})

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	err = er.ethclient.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Println(err)
		return "", nil, err
	}

	receipt, err := bind.WaitMined(ctx, er.ethclient, signedTx)
	if err != nil {
		return "", nil, err
	}

	if receipt.Status == 2 {
		return signedTx.Hash().Hex(), receipt, errors.New("transaction failed")
	}

	return signedTx.Hash().Hex(), receipt, nil
}
