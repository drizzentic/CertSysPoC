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

//Create student profile
func CreateStudentProfile(resp http.ResponseWriter, req *http.Request) {
	var s Student
	//Extract data from requests
	body, _ := ioutil.ReadAll(req.Body)

	json.Unmarshal(body, &s)

	//Extract the name and admission number to be used to create a unique file for student

	name := s.Name

	admissionNumber := s.AdmissionNumber
	filename := name + admissionNumber

	//Generate hash and create directory within the institution's folder

	directoryHash := utils.GetHash(filename)

	parentDir := utils.GetConfigs().Institution
	path := []string{}
	path = append(path, parentDir)
	path = append(path, string(directoryHash))
	newDir := strings.Join(path, "/")

	address := requestAddress(string(directoryHash))

	//Create a university account
	requestAddress(string(directoryHash))

	s.ResultsAddress = address

	utils.CreateDirIfNotExist(newDir)

	// Create transaction on the Blockchain
	response := createSimpleSpend(address, string(directoryHash))

	resp.Header().Set("Content-Type", "application/json")

	json.NewEncoder(resp).Encode(response)

	//Run operations to push the data to the ipfs nodes.
	go func() {
		//TODO: Add function to upload profile to ipfs directory
		//Recursively create the directory on the ipfs system

		//utils.DeleteDirIfExist(strings.Join(path,"/"))
	}()
}
func CreateUniversityProfile(university string) {

	//Generate hash and create directory within the institution's folder

	h := sha1.New()

	h.Write([]byte(university))

	universityHash := hex.EncodeToString(h.Sum(nil))

	requestAddress(string(universityHash))

}
