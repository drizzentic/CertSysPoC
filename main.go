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
	directory := utils.GetConnectionCredentials().Institution

	utils.CreateDirIfNotExist(directory)

}
func main() {

	//TODO:Add routes
	r := mux.NewRouter()
	r.HandleFunc("/get_address", controllers.RequestAddress).Methods(http.MethodGet)
	r.HandleFunc("/create_profile", controllers.CreateProfile)
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
