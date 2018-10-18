package controllers

import (
	"certSys/utils"
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

const (
	createAddress         = "getnewaddress"
	sendToAddress         = "sendtoaddress"
	getAccountAddress     = "getaccountaddress"
	getAddressesbyAccount = "getaddressesbyaccount"
	getAccountBalance     = "getbalance"
)

func requestAddress(filename string) string {
	var bcAddress string
	//Check that the user has no account

	params := []string{filename}

	address := processRpcCalls(getAddressesbyAccount, params)

	bcAddress = CheckResultsType(address)

	if bcAddress == "" {
		log.Println("Creating a new address for the student")
		newAddress := processRpcCalls(createAddress, params)

		return newAddress.Results.(string)

	}

	log.Println("Address Exists ", address.Results)
	return bcAddress
}

//Create transaction

func createSimpleSpend(address string, comment string) string {

	log.Println("This is the address", address)

	params := []string{address, utils.GetConfigs().TransactionValue, comment}

	txid := processRpcCalls(sendToAddress, params)

	//ToDO: Return appropriate error messages to the user
	if txid.Results == nil {

		return txid.Error.Message

	}
	return txid.Results.(string)
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
func createRawTransaction() {

}
func signTransaction() {

}

func sendTransaction() {

}

func SearchAccountbyHash() {
	//Returns the account name for the user based on the hash generated from name and admission

}
func checkAccountAddress() {

}
