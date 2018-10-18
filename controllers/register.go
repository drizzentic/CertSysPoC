package controllers

import (
	"certSys/utils"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type Student struct {
	Name            string `json:"fullname"`
	AdmissionNumber string `json:"admission_number"`
	School          string `json:"school"`
	Award           string `json:"award"`
	Department      string `json:"department"`
	ResultsAddress  string `json:"results_address"`
}

//Create student profile
func CreateProfile(resp http.ResponseWriter, req *http.Request) {
	var s Student

	//Extract data from requests
	body, _ := ioutil.ReadAll(req.Body)

	json.Unmarshal(body, &s)

	//Extract the name and admission number to be used to create a unique file for student

	name := s.Name

	admissionNumber := s.AdmissionNumber
	filename := name + admissionNumber

	//Generate hash and create directory within the institution's folder

	h := sha1.New()
	h.Write([]byte(filename))

	directoryHash := hex.EncodeToString(h.Sum(nil))

	parentDir := utils.GetConfigs().Institution
	path := []string{}
	path = append(path, parentDir)
	path = append(path, string(directoryHash))
	newDir := strings.Join(path, "/")

	address := requestAddress(string(directoryHash))
	s.ResultsAddress = address

	utils.CreateDirIfNotExist(newDir)

	//create a transaction on blockchain with the filename

	//Check balance

	createSimpleSpend(address, string(directoryHash))

	resp.Header().Set("Content-Type", "application/json")

	json.NewEncoder(resp).Encode(&s)

	//Run operations to push the data to the ipfs nodes.
	go func() {
		//TODO: Add function to upload profile to ipfs directory
		//Recursively create the directory on the ipfs system

		//utils.DeleteDirIfExist(strings.Join(path,"/"))
	}()
}