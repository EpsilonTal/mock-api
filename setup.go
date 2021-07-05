package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"github.tools.sap/atom-cfs/mock-api/utils"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func setup(w http.ResponseWriter, r *http.Request) {
	validateRequest(w, r)
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
	case "PATCH":
		updateTest(w, r)
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

func validateRequest(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		utils.GenerateResponse(w, 400, nil)
	}

	var input interface{}
	err := json.NewDecoder(r.Body).Decode(input)
	switch {
	case err == io.EOF:
		// empty body
		//utils.GenerateResponse(w, 400, nil)
	case err != nil:
		// other error
		utils.GenerateResponse(w, 400, nil)
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

	reqPostSettings, err := setMethodSettings(b, utils.PostKey)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	reqDeleteSettings, err := setMethodSettings(b, utils.DeleteKey)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	reqPatchSettings, err := setMethodSettings(b, utils.PatchKey)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	settings := &utils.MockConfig{
		ID:     testUUID,
		Post:   reqPostSettings,
		Delete: reqDeleteSettings,
		Patch:  reqPatchSettings,
	}
	mutex.Lock()
	tests[testUUID] = settings
	mutex.Unlock()
	log.Printf("Creating a test with ID: %s", testUUID)
	utils.GenerateResponse(w, 201, tests[testUUID])
}

func setMethodSettings(b []byte, key string) (utils.MethodConfig, error) {
	value := []byte(gjson.GetBytes(b, key).Raw)
	// Unmarshal
	var methodConfig utils.MethodConfig
	err := json.Unmarshal(value, &methodConfig)
	if err != nil {
		return utils.MethodConfig{}, err
	}
	return methodConfig, err
}

func updateTest(w http.ResponseWriter, r *http.Request) {
	testUUID := r.URL.Query().Get(utils.IDKey)

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

	reqPostSettings, err := setMethodSettings(b, utils.PostKey)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	reqDeleteSettings, err := setMethodSettings(b, utils.DeleteKey)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	reqPatchSettings, err := setMethodSettings(b, utils.PatchKey)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	settings := &utils.MockConfig{
		ID:     testUUID,
		Post:   reqPostSettings,
		Delete: reqDeleteSettings,
		Patch:  reqPatchSettings,
	}
	mutex.Lock()
	tests[testUUID] = settings
	mutex.Unlock()
	log.Printf("Updating a test with ID: %s", testUUID)
	utils.GenerateResponse(w, 200, tests[testUUID])
}

func getSingleTest(w http.ResponseWriter, r *http.Request) {
	testUUID := r.URL.Query().Get(utils.IDKey)
	if tests[testUUID] == nil {
		utils.GenerateResponse(w, 404, nil)
	} else {
		utils.GenerateResponse(w, 200, tests[testUUID])
	}
}
