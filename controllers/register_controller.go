package controllers

import (
	"certSys/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	Results string `json:"result"`
	Error   string `json:"error"`
	Id      string `json:"id"`
}

func Register() string {
	var address Response
	method := utils.Requests{createAddress}
	resp := utils.RpcCalls(&method)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(body)
	if err := json.Unmarshal(body, &address); err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return address.Results
}

//Create student profile
func (s *Student) CreateProfile(resp http.ResponseWriter, req *http.Request)  {

	//Extract data from requests

 s.ResultsAddress=Register()

 json.Marshal(s)

}
