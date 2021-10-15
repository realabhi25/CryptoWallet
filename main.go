package main

import (
	structures "CryptoWallet/Structures"
	"CryptoWallet/accountmanagement"
	"CryptoWallet/errorhandler"
	"CryptoWallet/wallet"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/golang/gddo/httputil/header"
	"github.com/gorilla/mux"
)

func createNewWallet(w http.ResponseWriter, r *http.Request) {

	wallet := wallet.NewWallet()
	newwallet := structures.WalletResponse{WalletAddress: wallet}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	b, _ := json.Marshal(newwallet)
	w.Write(b)
}

func getAllAccounts(w http.ResponseWriter, r *http.Request) {

	allaccounts := accountmanagement.GetAllAccounts()
	allaccountsres := structures.AllAccountsResponse{WalletAddresses: allaccounts}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	b, _ := json.Marshal(allaccountsres)
	w.Write(b)

}

func createNewAccount(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	w.Header().Set("Content-Type", "application/json")

	var newaccountrequest structures.NewAccountRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&newaccountrequest)

	if err != nil {

		msg, status := errorhandler.HandleHttpError(err, decoder)
		httperrorresponse := structures.GeneralHttpError{Message: msg}
		b, _ := json.Marshal(httperrorresponse)
		w.WriteHeader(status)
		w.Write(b)

	} else {

		newaccount := accountmanagement.CreateNewAccount(newaccountrequest.Password)
		newaccountresponse := structures.NewAccountResponse{WalletAddress: newaccount.String()}

		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(newaccountresponse)
		w.Write(b)
	}
}

func exportAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var exportAccountRequest structures.ExportAccountRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&exportAccountRequest)

	if err != nil {

		msg, status := errorhandler.HandleHttpError(err, decoder)
		httperrorresponse := structures.GeneralHttpError{Message: msg}
		b, _ := json.Marshal(httperrorresponse)
		w.WriteHeader(status)
		w.Write(b)

	} else {
		exportedaccount, err := accountmanagement.ExportAccount(exportAccountRequest.WalletAddress, exportAccountRequest.WalletPassword)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			httperrorresponse := structures.GeneralHttpError{Message: err.Error()}
			b, _ := json.Marshal(httperrorresponse)
			w.Write(b)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(exportedaccount)
		}
	}
}

func importAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	b, err := io.ReadAll(r.Body)
	println("imported : ", b)
	importedaccount, err := accountmanagement.ImportAccount(b, "abcd")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		httperrorresponse := structures.GeneralHttpError{Message: err.Error()}
		b, _ := json.Marshal(httperrorresponse)
		w.Write(b)
	} else {
		println(importedaccount.Address.String())
		importedaccountresponse := structures.ImportAccountResponse{Data: importedaccount}
		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(importedaccountresponse)
		w.Write(b)
	}
}

func main() {
	fmt.Println("Inside router")
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/wallet", createNewWallet).Methods(http.MethodPost)
	api.HandleFunc("/getaccounts", getAllAccounts).Methods(http.MethodPost)
	api.HandleFunc("/newaccount", createNewAccount).Methods(http.MethodPost)
	api.HandleFunc("/export", exportAccount).Methods(http.MethodPost)
	api.HandleFunc("/import", importAccount).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8081", r))
}
