package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//smtesting "github.tools.sap/atom-cfs/service-management/testing"

)

func main(){
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":1111",nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Hello")
	case "POST":
		callbackProvision(w,r)
		fmt.Printf("done")
	case "DELETE":
		callbackDeprovision(r)
	case "DEFAULT":
		fmt.Fprintf(w, "default")
	}
	return
}

func callbackProvision(w http.ResponseWriter, r *http.Request) *http.Response {
	w.Header().Set("Content-Type", "application/json")
	var message *string = nil
	//var newsList = make([]News, 0)
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil{
	return &http.Response{
		Status: "Bad request",
		StatusCode: 400,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body: nil,
		ContentLength: 0,
		Header: make(http.Header, 0),
	}}

	return &http.Response{
		Status: "200 OK",
		StatusCode: 200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body: nil,
		ContentLength: 0,
		Header: make(http.Header, 0),
	}
}

func callbackDeprovision(r *http.Request) *http.Response {
	name := r.URL.Query().Get("name")
	namespace := r.URL.Query().Get("namespace")
	uid := r.URL.Query().Get("uid")
	if len(name) < 1 || len(namespace) < 1 || len(uid) == 0{
		return &http.Response{
			Status: "Bad request",
			StatusCode: 400,
			Proto:         "HTTP/1.1",
			ProtoMajor:    1,
			ProtoMinor:    1,
			Body: nil,
			ContentLength: 0,
			Header: make(http.Header, 0),
		}
		//need to delete the resource instance by the name
	}


	return nil
}