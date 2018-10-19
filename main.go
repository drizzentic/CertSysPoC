package main

import (
	"certSys/controllers"
	"certSys/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func init() {

	//Create ipfs folder for the university
	directory := utils.GetConfigs().Institution
	//Create Account for the university
	controllers.CreateUniversityProfile(directory)

	//Create directory
	utils.CreateDirIfNotExist(directory)

}
func main() {

	//TODO:Add routes
	r := mux.NewRouter()
	r.HandleFunc("/results", controllers.GetResults).Methods(http.MethodPost)
	r.HandleFunc("/results/create", controllers.CreateResults).Methods(http.MethodPost)
	r.HandleFunc("/profile", controllers.CreateStudentProfile)
	http.Handle("/", r)

	//TODO: Initialize API server
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
