package controllers

import (
	"certSys/utils"
	"fmt"
	"log"
)

type Response struct {
	Results interface{} `json:"result"`
	Error   `json:"error"`
	Id      string `json:"id"`
}
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type rawTransaction struct {
	Txid string `json:"txid"`
	Vout int    `json:"vout"`
}

const (
	createAddress         = "getnewaddress"
	sendToAddress         = "sendtoaddress"
	getAccountAddress     = "getaccountaddress"
	getAddressesbyAccount = "getaddressesbyaccount"
	getAccountBalance     = "getbalance"
	generate              = "generate"
	createRawTransaction  = "createrawtransaction"
)

func requestAddress(filename string) string {
	var bcAddress string
	//Check that the user has no account

	params := []string{filename}

	address := processRpcCalls(getAddressesbyAccount, params, "")

	bcAddress = CheckResultsType(address)

	if bcAddress == "" {
		log.Println("Creating a new address for the student")
		newAddress := processRpcCalls(createAddress, params, "")

		return newAddress.Results.(string)

	}

	log.Println("Address Exists ", address.Results)
	return bcAddress
}

//Create transaction

func createSimpleSpend(address string, comment string) Response {

	log.Println("This is the address", address)

	params := []string{address, utils.GetConfigs().TransactionValue, comment}

	txid := processRpcCalls(sendToAddress, params, "")
	//TODO: Return appropriate error messages to the user
	if (Error{}) != txid.Error {
		errorCode := txid.Error.Code
		if errorCode != 0 {
			//Invoke generate new blocks to add balance
			processRpcCalls(generate, []string{"1"}, "")
		}
		return txid

	}
	return txid
}

func CheckResultsType(a Response) string {
	var bcAddress string

	str, ok := a.Results.(string)
	if ok {

		bcAddress = str

	} else {
		for _, v := range a.Results.([]interface{}) {

			bcAddress = v.(string)
		}

	}

	return bcAddress
}
func checkAccountAddress(accountName string) Response {

	address := processRpcCalls(getAccountAddress, []string{accountName}, "")

	return address

}
func createTransaction(txid string, data string) {
	raw := []string{txid}
	result := processRpcCalls(createRawTransaction, raw, data)

	//Sign transaction
	fmt.Print(result)

	signTransaction(result)

}
func signTransaction(r Response) {

}

func sendTransaction() {

}

func SearchAccountbyHash() {
	//Returns the account name for the user based on the hash generated from name and admission

}
