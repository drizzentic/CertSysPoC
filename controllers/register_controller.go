package controllers


import (
"certSys/utils"
"encoding/json"
"io/ioutil"
"log"
)

const (
	createAddress = "getnewaddress"
)

type Student struct {
	Name            string `json:"fullname"`
	AdmissionNumber string `json:"admission_number"`
	School          string `json:"school"`
	Award           string `json:"award"`
	Department      string `json:"department"`
	ResultsAddress  string `json:"results_address"`
}
type Response struct {
	Results string	`json:"result"`
	Error   string `json:"error"`
	Id      string `json:"id"`
}
func Register(s *Student) string {
	var address Response
	method := utils.Requests{createAddress}
	resp := utils.RpcCalls(&method)
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &address); err != nil {
		panic(err)
	}
	log.Println("The new address is: ",address.Results)
	defer resp.Body.Close()
	return address.Results
}
