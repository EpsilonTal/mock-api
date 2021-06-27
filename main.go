package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Printf("starting mock server")
	http.HandleFunc("/v1/provision", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fail := r.URL.Query().Get("fail")

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Fake Callback Application")
	case "POST":
		handleProvision(w, fail)
	case "DELETE":
		handleDeprovision(w, fail)
	}
	return
}

func handleProvision(w http.ResponseWriter, fail string) {
	if fail == "true" {
		w = generateResponse(w, 400, nil)
		return
	}

	body := []byte("{\"accessUrl\": \"accessUrl.com\",\"metadata\": {\"key1\": \"value1\",\"key2\": \"value2\"}}")
	w = generateResponse(w, 201, body)
}

func handleDeprovision(w http.ResponseWriter, fail string) {
	if fail == "true" {
		w = generateResponse(w, 410, nil)
		return
	}

	w = generateResponse(w, 200, nil)
}

func generateResponse(w http.ResponseWriter, statusCode int, body []byte) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(body)
	return w
}
