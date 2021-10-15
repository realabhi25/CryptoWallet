package accountmanagement

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

func getKeyStore() *keystore.KeyStore {

	ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)

	// Create a new account with the specified encryption passphrase.
	// newAcc, _ := ks.NewAccount("Creation password")
	// fmt.Println(newAcc)

	return ks

	// // Export the newly created account with a different passphrase. The returned
	// // data from this method invocation is a JSON encoded, encrypted key-file.
	// jsonAcc, _ := ks.Export(newAcc, "Creation password", "Export password")

	// // Update the passphrase on the account created above inside the local keystore.
	// _ = ks.Update(newAcc, "Creation password", "Update password")

	// // Delete the account updated above from the local keystore.
	// _ = ks.Delete(newAcc, "Update password")

	// // Import back the account we've exported (and then deleted) above with yet
	// // again a fresh passphrase.
	// impAcc, _ := ks.Import(jsonAcc, "Export password", "Import password")

}
func getAccountManager(ks *keystore.KeyStore) *accounts.Manager {
	am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
	return am
}
func GetAllAccounts() []common.Address {

	ks := getKeyStore()
	am := getAccountManager(ks)
	return am.Accounts()
}

func CreateNewAccount(password string) common.Address {
	ks := getKeyStore()
	newAcc, _ := ks.NewAccount(password)
	return newAcc.Address
}

func ExportAccount(address common.Address, password string) ([]byte, error) {
	ks := getKeyStore()
	account := accounts.Account{Address: address}
	jsonAcc, err := ks.Export(account, password, password)
	println("exported", jsonAcc)
	return jsonAcc, err
}

func ImportAccount(keyjson []byte, password string) (accounts.Account, error) {
	ks := getKeyStore()
	impAcc, err := ks.Import(keyjson, password, password)
	return impAcc, err
}
