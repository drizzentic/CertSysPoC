package controllers

import (
	"certSys/utils"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

func RequestAddress(resp http.ResponseWriter, req *http.Request) {
	var address Response
	method := utils.Requests{createAddress}
	response := utils.RpcCalls(&method)
	body, _ := ioutil.ReadAll(response.Body)
	if err := json.Unmarshal(body, &address); err != nil {
		panic(err)
	}
	defer response.Body.Close()

	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(&address)
}

//Create student profile
func CreateProfile(resp http.ResponseWriter, req *http.Request) {
	var s Student
	//Extract data from requests
	body, _ := ioutil.ReadAll(req.Body)

	json.Unmarshal(body, &s)

	//Extract the name and admission number to be used to create a unique file for student

	name := s.Name
	address := s.ResultsAddress
	admissionNumber := s.AdmissionNumber

	filename := name + address + admissionNumber

	//Generate hash and create directory within the institution's folder

	h := sha1.New()
	h.Write([]byte(filename))

	directoryHash := hex.EncodeToString(h.Sum(nil))

	parentDir := utils.GetConnectionCredentials().Institution
	path := []string{}
	path = append(path, parentDir)
	path = append(path, string(directoryHash))
	newDir := strings.Join(path, "/")

	fmt.Print(string(newDir))

	utils.CreateDirIfNotExist(newDir)

	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(&s)

	//Run operations to push the data to the ipfs nodes.
	go func() {
		//TODO: Add function to upload profile to ipfs directory
		//Recursively create the directory on the ipfs system

		//utils.DeleteDirIfExist(strings.Join(path,"/"))
	}()
}
