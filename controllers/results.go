package controllers

import (
	"certSys/utils"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	s Student
	o Output
	r Response
)

func GetResults(r http.ResponseWriter, w *http.Request) {
	body, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal(body, &s)

	folderDirectory := utils.GetHash(s.Name + s.AdmissionNumber)

	//List transactions by account

	transactions := processRpcCalls(listTransactions, []string{folderDirectory}, "")

	r.Header().Set("Content-Type", "application/json")
	json.NewEncoder(r).Encode(transactions.Results)
}
func CreateResults(r http.ResponseWriter, w *http.Request) {
	body, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal(body, &o)
	//TODO:Create json file for the data and post hash as transaction to blockchain
	//Get directory for the user
	fileFolder := utils.GetHash(o.OverallResults.Name + o.OverallResults.AdmissionNumber)

	filesDirectory := utils.GetConfigs().Institution + "/" + fileFolder

	//Create file name
	h := sha1.New()
	h.Write([]byte(body))
	a := fmt.Sprintf("%x", h.Sum(nil))

	//TODO: before creating a transaction, validate that the file doesn't exist
	response := createSimpleSpend(checkAccountAddress(fileFolder).Results.(string), a)

	ioutil.WriteFile(filesDirectory+"/"+a+".json", body, 0644)

	r.Header().Set("Content-Type", "application/json")

	json.NewEncoder(r).Encode(response)

}
