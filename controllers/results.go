package controllers

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/CertSysPoC/utils"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	s Student
	o Output
	r Response
)

func GetResults(r http.ResponseWriter, w *http.Request) {
	body, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal(body, &s)
	fmt.Print(s.Name)
	folderDirectory := utils.GetHash(s.Name + s.AdmissionNumber)

	//List transactions by account

	transactions := processRpcCalls(listTransactions, []string{folderDirectory}, "")

	r.Header().Set("Content-Type", "application/json")
	json.NewEncoder(r).Encode(transactions.Results)
}

func GetTranscript(r http.ResponseWriter, w *http.Request) {

	filename := w.URL.Query().Get("hash")
	student := w.URL.Query().Get("student")
	institution := utils.GetConfigs().Institution
	jsonFile, err := os.Open(institution + "/" + student + "/" + filename + ".json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	r.Header().Set("Content-Type", "application/json")
	json.NewEncoder(r).Encode(jsonFile)
}

func CreateResults(r http.ResponseWriter, w *http.Request) {
	body, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal(body, &o)

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
