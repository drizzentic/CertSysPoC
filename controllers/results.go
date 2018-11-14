package controllers

import (
	"certSys/utils"
	"crypto/sha256"
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
	filesDirectory ,folderDirectory := utils.GetDirectory(s.Name , s.AdmissionNumber)

	//List transactions by account

	transactions := processRpcCalls(listTransactions, []string{folderDirectory}, "")

	utils.HttpCalls(true,filesDirectory+"/05fd7777451f071e640b91eee2db27802af8c914da7fbccaa6e559ecb3e25aa4.json")
	r.Header().Set("Content-Type", "application/json")
	json.NewEncoder(r).Encode(transactions.Results)
}
func CreateResults(r http.ResponseWriter, w *http.Request) {
	body, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal(body, &o)

	filesDirectory,fileFolder:=utils.GetDirectory(o.OverallResults.Name,o.OverallResults.AdmissionNumber)

	//Create file name
	h := sha256.New()
	h.Write([]byte(body))
	a := fmt.Sprintf("%x", h.Sum(nil))

	//TODO: before creating a transaction, validate that the file doesn't exist
	response := createSimpleSpend(checkAccountAddress(fileFolder).Results.(string), a)

	ioutil.WriteFile(filesDirectory+"/"+a+".json", body, 0644)

	r.Header().Set("Content-Type", "application/json")

	json.NewEncoder(r).Encode(response)

}
