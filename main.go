package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	//"sync"
)

var (
	tests map[string]testSettings
	mutex *sync.Mutex
)

type testSettings struct {
	ID                    string      `json:"id"`
	ProvisionStatus       int64       `json:"provisionStatus"`
	ProvisionResponseBody interface{} `json:"provisionResponseBody"`
	DeprovisionStatus     int64       `json:"deprovisionStatus"`
}

func main() {
	tests = make(map[string]testSettings)
	mutex = &sync.Mutex{}
	log.Printf("starting mock server")

	/*
		POST:
		 Create a new test configuration, using testSettings struct
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

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Fake Callback Application")
	case "POST":
		handleProvision(w, r)
	case "DELETE":
		handleDeprovision(w, r)
	}
	return
}
func setup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTest(w, r)
		return
	case "POST":
		createTest(w, r)
		return
	case "DELETE":
		deleteTest(w, r)
		return
	}
}

func createTest(w http.ResponseWriter, r *http.Request) {
	testUUID := uuid.New().String()

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var reqTestSettings testSettings
	err = json.Unmarshal(b, &reqTestSettings)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	mutex.Lock()
	settings := testSettings{
		ID:                    testUUID,
		ProvisionStatus:       gjson.GetBytes(b, "provisionStatus").Int(),
		ProvisionResponseBody: gjson.GetBytes(b, "provisionResponseBody").Value(),
		DeprovisionStatus:     gjson.GetBytes(b, "deprovisionStatus").Int(),
	}
	tests[testUUID] = settings
	mutex.Unlock()
	log.Printf("Creating a test with ID: %s", testUUID)
	generateResponse(w, 201, tests[testUUID])
}

func getTest(w http.ResponseWriter, r *http.Request) {
	testUUID := r.URL.Query().Get("id")
	generateResponse(w, 200, tests[testUUID])
}

func deleteTest(w http.ResponseWriter, r *http.Request) {
	deleteAll := r.URL.Query().Get("all")
	if deleteAll == "true" {
		tests = make(map[string]testSettings)
	} else {
		testUUID := r.URL.Query().Get("id")
		delete(tests, testUUID)
	}
	generateResponse(w, 200, nil)
}

func handleProvision(w http.ResponseWriter, r *http.Request) {
	testUUID := r.URL.Query().Get("id")
	w = generateResponse(w, int(tests[testUUID].ProvisionStatus), tests[testUUID].ProvisionResponseBody)
}

func handleDeprovision(w http.ResponseWriter, r *http.Request) {
	testUUID := r.URL.Query().Get("id")
	w = generateResponse(w, int(tests[testUUID].DeprovisionStatus), nil)
}

func generateResponse(w http.ResponseWriter, statusCode int, object interface{}) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	marshal, _ := json.Marshal(object)
	w.Write(marshal)
	return w
}
