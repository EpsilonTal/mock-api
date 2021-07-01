package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"github.tools.sap/atom-cfs/mock-api/utils"
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
		getAll := r.URL.Query().Get(utils.AllKey)
		if getAll == "true" {
			utils.GenerateResponse(w, 200, tests)
		} else {
			getSingleTest(w, r)
		}
		return
	case "DELETE":
		deleteAll := r.URL.Query().Get(utils.AllKey)
		if deleteAll == "true" {
			tests = make(map[string]*utils.MockConfig)
		} else {
			testUUID := r.URL.Query().Get(utils.IDKey)
			delete(tests, testUUID)
		}
		utils.GenerateResponse(w, 200, nil)
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
	var reqTestSettings utils.MockConfig
	err = json.Unmarshal(b, &reqTestSettings)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	mutex.Lock()
	settings := &utils.MockConfig{
		ID:                   testUUID,
		PostResponseStatus:   gjson.GetBytes(b, utils.PostResponseStatusKey).Int(),
		PostResponseBody:     gjson.GetBytes(b, utils.PostResponseBodyKey).Value(),
		DeleteResponseStatus: gjson.GetBytes(b, utils.DeleteResponseStatusKey).Int(),
	}
	tests[testUUID] = settings
	mutex.Unlock()
	log.Printf("Creating a test with ID: %s", testUUID)
	utils.GenerateResponse(w, 201, tests[testUUID])
}

func getSingleTest(w http.ResponseWriter, r *http.Request) {
	testUUID := r.URL.Query().Get(utils.IDKey)
	if tests[testUUID] == nil {
		utils.GenerateResponse(w, 404, nil)
	} else {
		utils.GenerateResponse(w, 200, tests[testUUID])
	}
}
