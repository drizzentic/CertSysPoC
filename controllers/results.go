package controllers

import "net/http"

type Course struct {
	Code         string `json:"code"`
	Description  string `json:"description"`
	CF           int    `json:"cf"`
	Grade        string `json:"grade"`
	AcademicYear string `json:"academic_year"`
	Results      `json:"results"`
}

type Results struct {
	Averages       float64 `json:"averages"`
	Cumulative     float64 `json:"cumulative"`
	Recommendation string  `json:"recommendation"`
}

func GetResults(r http.ResponseWriter, w *http.Request) {

}
func CreateResults(r http.ResponseWriter, w *http.Request) {

}
