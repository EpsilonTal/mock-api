package main

import (
	"fmt"
	"github.tools.sap/atom-cfs/mock-api/utils"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		handlePost(w, r)
	case "GET":
		fmt.Fprintf(w, "Fake Callback Application")
	case "DELETE":
		handleDelete(w, r)
	}
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	testUUID := r.URL.Query().Get(utils.IDKey)
	w = utils.GenerateResponse(w, int(tests[testUUID].PostResponseStatus), tests[testUUID].PostResponseBody)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	testUUID := r.URL.Query().Get(utils.IDKey)
	w = utils.GenerateResponse(w, int(tests[testUUID].DeleteResponseStatus), nil)
}
