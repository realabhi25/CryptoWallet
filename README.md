# CryptoWallet

## An easy to setup local crypto wallet based on [Geth](https://geth.ethereum.org) 

### To run.
```
go run CrytoWallet
```

This will expose a set a api's.

### To Create new Wallet
```
curl --location --request POST 'http://localhost:8081/api/v1/newaccount' \
--header 'Content-Type: application/json' \
--data-raw '{
    "WalletPassword" : "<wallet password>"
}'
```

### To Get all the created wallets
```
curl --location --request POST 'http://localhost:8081/api/v1/getaccounts'
```

### To Export all the created addresses to another wallet / crypto platform
```
curl --location --request POST 'http://localhost:8081/api/v1/export' \
--header 'Content-Type: application/json' \
--data-raw '{
    "WalletAddress" : "<wallet address>",
    "WalletPassword" : "<wallet password>"
}'
```

#### You can connect this project with Ethereum mainnet by replacing the https://github.com/realabhi25/CryptoWallet/blob/main/network/main.go with your Infura API Key