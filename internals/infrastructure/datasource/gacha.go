package datasource

import (
	"os"
	"time"

	"github.com/jinzhu/gorm"
	godotenv "github.com/joho/godotenv"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"

	"context"
	"fmt"
	"log"

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

func (gr *gachaRepository) TransferToken() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Failed to read file: ", err)
	}
	client, err := ethclient.Dial("https://eth-ropsten.alchemyapi.io/v2/vxAxsV1d-Ry0cxaQgEu6g-c1D5U0sbdm")
	if err != nil {
		log.Fatal(err)
	}

	// load private key of the Wallet
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()                   // 公開鍵を含むインタフェースをreturn
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) // 型アサーション、publicKey変数の型を明示的に設定
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Println("fromAddress: ", fromAddress)

	// 次のトランザクションに使用するnonceの読み込み
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// トランザクションのETHの量の設定、ERC20を転送するためETHの値は０。Tokenの値はdataのフィールドに設定する
	// convert ETH to wei
	// 18 decimal places, 1ETH = 1000000000000000000(1 + 18 zeros)
	// Token transfers don't require ETH to be transferred so set the value to 0
	value := big.NewInt(0) // = wei (0 eth)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Wallet address 0x は１０進数を１６進数で表している
	toAddress := common.HexToAddress("0xEa58D2fFBa020c4f3152dB37E8896B4d233F849b")
	// Token contract address
	tokenAddress := common.HexToAddress("0x984a6eaecBE9e77339931E191B6bf314c6f65dab")

	transferFnSignature := []byte("transfer(address,uint256)")
	// Get the method ID of the function
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4] // The first 4 bytes of the resulting hash is the methodId: コントラクトのメソッドをbyte形式にしてKECCAK-256でハッシュ化し、先頭から４バイト取ってきたもの
	fmt.Printf("Method ID: %s\n", hexutil.Encode(methodID))

	// zero pad (to the left) the account address. The resulting byte slice must be 32 bytes long.
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Printf("To address: %s\n", hexutil.Encode(paddedAddress))

	amount := new(big.Int)
	amount.SetString("1000000000000000000", 10) // 1 token
	// zero pad (to the left) the amount. The resulting byte slice must be 32 bytes long.
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Printf("Token amount: %s", hexutil.Encode(paddedAmount))

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	estimatedGas, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}

	gasLimit := uint64(float64(estimatedGas) * 1.80)

	log.Println("Gas Limit:", gasLimit)
	// Transaction
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	// sign the transaction with the private key of the sender
	// The SignTx method requires the EIP155 signer.
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// broadcast the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tokens sent at TX: %s", signedTx.Hash().Hex())

	return err
}
