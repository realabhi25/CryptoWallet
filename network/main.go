package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	_, err := ethclient.Dial("https://mainnet.infura.io/v3/481aae13af304546be954575713bf8c6")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
}
