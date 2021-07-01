package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Fake Callback Application")
	case "POST":
		handlePost(w, r)
	case "DELETE":
		handleDelete(w, r)
	}
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	testUUID := r.URL.Query().Get(idKey)
	w = generateResponse(w, int(tests[testUUID].PostResponseStatus), tests[testUUID].PostResponseBody)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	testUUID := r.URL.Query().Get(idKey)
	w = generateResponse(w, int(tests[testUUID].DeleteResponseStatus), nil)
}
