package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
)

func setup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		createTest(w, r)
		return
	case "GET":
		getTest(w, r)
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
	var reqTestSettings MockConfig
	err = json.Unmarshal(b, &reqTestSettings)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	mutex.Lock()
	settings := &MockConfig{
		ID:                   testUUID,
		PostResponseStatus:   gjson.GetBytes(b, PostResponseStatusKey).Int(),
		PostResponseBody:     gjson.GetBytes(b, PostResponseBodyKey).Value(),
		DeleteResponseStatus: gjson.GetBytes(b, DeleteResponseStatusKey).Int(),
	}
	tests[testUUID] = settings
	mutex.Unlock()
	log.Printf("Creating a test with ID: %s", testUUID)
	generateResponse(w, 201, tests[testUUID])
}

func getTest(w http.ResponseWriter, r *http.Request) {
	testUUID := r.URL.Query().Get(idKey)
	if tests[testUUID] == nil {
		generateResponse(w, 404, nil)
	} else {
		generateResponse(w, 200, tests[testUUID])
	}
}

func deleteTest(w http.ResponseWriter, r *http.Request) {
	deleteAll := r.URL.Query().Get(deleteAllKey)
	if deleteAll == "true" {
		tests = make(map[string]*MockConfig)
	} else {
		testUUID := r.URL.Query().Get(idKey)
		delete(tests, testUUID)
	}
	generateResponse(w, 200, nil)
}
