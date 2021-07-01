package main

import (
	"encoding/json"
	"net/http"
)

const (
	PostResponseStatusKey   = "postResponseStatus"
	PostResponseBodyKey     = "postResponseBody"
	DeleteResponseStatusKey = "deleteResponseStatus"
	idKey                   = "id"
	deleteAllKey            = "all"
)

func generateResponse(w http.ResponseWriter, statusCode int, object interface{}) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	marshal, _ := json.Marshal(object)
	w.Write(marshal)
	return w
}
