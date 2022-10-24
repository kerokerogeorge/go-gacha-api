package datasource

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	godotenv "github.com/joho/godotenv"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"

	"fmt"

	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

type Gacha struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type gachaRepository struct {
	db *gorm.DB
}

func NewGachaRepository(database *gorm.DB) *gachaRepository {
	db := database
	return &gachaRepository{
		db: db,
	}
}

func (gr *gachaRepository) CreateGacha(gacha *model.Gacha) (*model.Gacha, error) {
	err := gr.db.Table("gachas").Create(gacha).Error
	if err != nil {
		return nil, err
	}
	return gacha, nil
}

func (gr *gachaRepository) List() ([]*model.Gacha, error) {
	var gachas []*model.Gacha
	err := gr.db.Find(&gachas).Error
	if err != nil {
		return nil, err
	}
	return gachas, nil
}

func (gr *gachaRepository) GetOne(gachaId string) (*model.Gacha, error) {
	var gacha Gacha
	err := gr.db.Table("gachas").Where("id = ?", gachaId).First(&gacha).Error
	if err != nil {
		return nil, err
	}
	return gr.ToGachaModel(gacha), nil
}

func (gr *gachaRepository) DeleteGacha(gacha *model.Gacha) error {
	err := gr.db.Delete(&gacha).Error
	if err != nil {
		return err
	}
	return nil
}

func (gr *gachaRepository) ToGachaModel(gacha Gacha) *model.Gacha {
	return &model.Gacha{
		ID:        gacha.ID,
		CreatedAt: gacha.CreatedAt,
		UpdatedAt: gacha.UpdatedAt,
	}
}

func (gr *gachaRepository) TransferToken(ctx *gin.Context) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		return "", err
	}
	client, err := ethclient.Dial(os.Getenv("URL"))
	if err != nil {
		log.Println(err)
		return "", err
	}

	// load private key of the Wallet
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Println(err)
		return "", err
	}

	publicKey := privateKey.Public()                   // 公開鍵を含むインタフェースをreturn
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) // 型アサーション、publicKey変数の型を明示的に設定
	if !ok {
		log.Println(err)
		return "", err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 次のトランザクションに使用するnonceの読み込み
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// トランザクションのETHの量の設定、ERC20を転送するためETHの値は０。Tokenの値はdataのフィールドに設定する
	// convert ETH to wei
	// 18 decimal places, 1ETH = 1000000000000000000(1 + 18 zeros)
	// Token transfers don't require ETH to be transferred so set the value to 0
	value := big.NewInt(0) // = wei (0 eth)

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Wallet address 0x は１０進数を１６進数で表している
	toAddress := common.HexToAddress(os.Getenv("TO_ADDRESS"))
	// Token contract address
	tokenAddress := common.HexToAddress(os.Getenv("TOKEN_ADDRESS"))

	transferFnSignature := []byte("transfer(address,uint256)")
	// Get the method ID of the function
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4] // The first 4 bytes of the resulting hash is the methodId: コントラクトのメソッドをbyte形式にしてKECCAK-256でハッシュ化し、先頭から４バイト取ってきたもの

	// zero pad (to the left) the account address. The resulting byte slice must be 32 bytes long.
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	amount := new(big.Int)
	amount.SetString("10000000000000000000", 10) // 10 token
	// zero pad (to the left) the amount. The resulting byte slice must be 32 bytes long.
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	estimatedGas, err := client.EstimateGas(ctx, ethereum.CallMsg{
		From:     fromAddress,
		To:       &tokenAddress,
		Data:     data,
		Value:    value,
		GasPrice: gasPrice,
	})
	if err != nil {
		log.Println(err)
		return "", err
	}

	gasLimit := uint64(float64(estimatedGas))
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &tokenAddress,
		Value:    value,
		Gas:      estimatedGas,
		GasPrice: gasPrice,
		Data:     data,
	})
	// sign the transaction with the private key of the sender
	// The SignTx method requires the EIP155 signer.
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Println(err)
		return "", err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// broadcast the transaction
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Println(err)
		return "", err
	}

	fmt.Printf("===================================================")
	fmt.Printf("nonce: %d\n", nonce)
	fmt.Printf("From Address: %s\n", fromAddress)
	fmt.Printf("TOKEN Address: %s\n", tokenAddress)
	fmt.Printf("To address: %s\n", hexutil.Encode(paddedAddress))
	fmt.Printf("Token amount: %s\n", hexutil.Encode(paddedAmount))
	fmt.Printf("Method ID: %s\n", hexutil.Encode(methodID))
	fmt.Printf("data length: %d\n", len(data))
	fmt.Printf("estimated Gas: %d\n", estimatedGas)
	fmt.Printf("Gas Limit: %d\n", gasLimit)
	fmt.Printf("Gas Price: %d\n", gasPrice)
	fmt.Printf("Tokens sent at TX: %s\n", signedTx.Hash().Hex())
	fmt.Printf("===================================================")

	return signedTx.Hash().Hex(), err
}
