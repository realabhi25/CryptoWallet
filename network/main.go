package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	_, err := ethclient.Dial("https://mainnet.infura.io/v3/<Your Infura Key>")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
}
