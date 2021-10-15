package structures

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
)

type NewAccountRequest struct {
	Password string `json:"WalletPassword,omitempty" bson:"WalletPassword,omitempty" binding:"required"`
}

type NewAccountResponse struct {
	WalletAddress string
}

type AllAccountsResponse struct {
	WalletAddresses []common.Address
}

type ExportAccountRequest struct {
	WalletAddress  common.Address `json:"WalletAddress,omitempty" bson:"WalletAddress,omitempty" binding:"required"`
	WalletPassword string         `json:"WalletPassword,omitempty" bson:"WalletPassword,omitempty" binding:"required"`
}

// type ExportAccountResponse struct {
// 	ExportedAccount []byte
// }

type ImportAccountRequest struct {
	KeyJson        string
	WalletPassword string
}

type ImportAccountResponse struct {
	Data accounts.Account
}
