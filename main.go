package main

import (
	"github.com/CertSysPoC/controllers"
	"github.com/CertSysPoC/utils"
	"github.com/gorilla/handlers"
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
	r.HandleFunc("/results", controllers.GetResults).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/results/create", controllers.CreateResults).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/profile", controllers.CreateStudentProfile).Methods(http.MethodPost, http.MethodOptions)
	http.Handle("/", r)

	//TODO: Initialize API server
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8001",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// start server listen
	// with error handling
	log.Fatal(srv.ListenAndServe(), handlers.CORS(originsOk, headersOk, methodsOk)(r))
	//log.Fatal(srv.ListenAndServe(":" + os.Getenv("PORT"), handlers.CORS(originsOk, headersOk, methodsOk)(r)))

}
