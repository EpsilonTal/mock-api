package main

import (
	"github.tools.sap/atom-cfs/mock-api-app/utils"
	"log"
	"net/http"
	"sync"
)

var (
	tests map[string]*utils.MockConfig
	mutex *sync.Mutex
)

func main() {
	tests = make(map[string]*utils.MockConfig)
	mutex = &sync.Mutex{}
	log.Printf("starting mock server")

	/*
		POST:
		 Create a new test configuration, using MockConfig struct
		 Response will contain the ID of the test config
		DELETE:
		 Optional flags:
			* all=true: delete all test configurations
			* id=uuid: delete single test configuration
	*/
	http.HandleFunc("/setup", setup)

	/*
		Mandatory flag:
		* id=testUUID: with define which test config to work with
		POST:
		 Execute create with the test config
		DELETE:
		 Execute delete with the test config
	*/
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
