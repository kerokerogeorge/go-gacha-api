package config

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

func NewEthClient() *ethclient.Client {
	ethClient, err := ethclient.Dial(os.Getenv("URL"))
	if err != nil {
		log.Fatal(err)
	}
	return ethClient
}
