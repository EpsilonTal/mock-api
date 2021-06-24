package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":1111", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fail := r.URL.Query().Get("fail")

	switch r.Method {
	case "GET":
		fmt.Print(w, "Fake Callback Application")
	case "POST":
		handleProvision(w, r, fail)
	case "DELETE":
		handleDeprovision(w, r, fail)
	}
	return
}

func handleProvision(w http.ResponseWriter, r *http.Request, fail string) {
	if fail == "true" {
		w = generateResponse(w, 400, nil)
		return
	}

	body := []byte("{\"accessUrl\": \"accessUrl.com\",\"metadata\": {\"key1\": \"value1\",\"key2\": \"value2\"}}")
	w = generateResponse(w, 201, body)
}

func handleDeprovision(w http.ResponseWriter, r *http.Request, fail string) {
	valid := validateDeprovisionRequest(r, fail)
	if !valid {
		w = generateResponse(w, 410, nil)
		return
	}

	w = generateResponse(w, 200, nil)
}

func validateDeprovisionRequest(r *http.Request, fail string) bool {
	name := r.URL.Query().Get("name")
	namespace := r.URL.Query().Get("namespace")
	uid := r.URL.Query().Get("uid")

	if fail == "true" || len(name) == 0 || len(namespace) == 0 || len(uid) == 0 {
		return false
	}
	return true
}

func generateResponse(w http.ResponseWriter, statusCode int, body []byte) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(body)
	return w
}
