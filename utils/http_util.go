package utils

import (
	"encoding/json"
	"net/http"
)

const (
	PostResponseStatusKey   = "postResponseStatus"
	PostResponseBodyKey     = "postResponseBody"
	DeleteResponseStatusKey = "deleteResponseStatus"
	IDKey                   = "id"
	AllKey                  = "all"
)

func GenerateResponse(w http.ResponseWriter, statusCode int, object interface{}) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	marshal, _ := json.Marshal(object)
	w.Write(marshal)
	return w
}
