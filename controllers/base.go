package controllers

import (
	"encoding/json"
	"github.com/CertSysPoC/utils"
	"io/ioutil"
)

type Output struct {
	Results []struct {
		Description string  `json:"description"`
		CF          float64 `json:"cf"`
		Grade       string  `json:"grade"`
		Name        string  `json:"fullname"`
		Code        string  `json:"code"`
	} `json:"results"`
	OverallResults struct {
		Name            string  `json:"fullname"`
		Code            string  `json:"code"`
		AcademicYear    string  `json:"academic_year"`
		Averages        float64 `json:"averages"`
		Cumulative      float64 `json:"cumulative"`
		Recommendation  string  `json:"recommendation"`
		AdmissionNumber string  `json:"admission_number"`
	} `json:"overalls"`
}
type Student struct {
	Name            string `json:"fullname"`
	AdmissionNumber string `json:"admission_number"`
	School          string `json:"school"`
	Award           string `json:"award"`
	Department      string `json:"department"`
	ResultsAddress  string `json:"results_address"`
}

func processRpcCalls(m string, params []string, data string) Response {
	var address Response

	method := utils.Requests{m, params}
	response := utils.RpcCalls(&method, params, data, 0)

	body, _ := ioutil.ReadAll(response.Body)

	if err := json.Unmarshal(body, &address); err != nil {

		panic(err)
	}

	defer response.Body.Close()

	return address
}
