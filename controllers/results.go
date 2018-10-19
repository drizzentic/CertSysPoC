package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Output struct {
	Results []struct {
		Code        string  `json:"code"`
		Description string  `json:"description"`
		CF          float64 `json:"cf"`
		Grade       string  `json:"grade"`
	} `json:"results"`
	OverallResults struct {
		AcademicYear   string  `json:"academic_year"`
		Averages       float64 `json:"averages"`
		Cumulative     float64 `json:"cumulative"`
		Recommendation string  `json:"recommendation"`
	} `json:"overalls"`
}

type OverallResults struct {
}

func GetResults(r http.ResponseWriter, w *http.Request) {

}
func CreateResults(r http.ResponseWriter, w *http.Request) {
	var s Output
	body, _ := ioutil.ReadAll(w.Body)
	//TODO:Create json file for the data and post hash as transaction to blockchain
	ioutil.WriteFile("output.json", body, 0644)
	json.Unmarshal(body, &s)

	r.Header().Set("Content-Type", "application/json")

	json.NewEncoder(r).Encode(s)

}
