package utils

type MockConfig struct {
	ID                   string      `json:"id"`
	PostResponseStatus   int64       `json:"postResponseStatus"`
	PostResponseBody     interface{} `json:"postResponseBody"`
	DeleteResponseStatus int64       `json:"deleteResponseStatus"`
}

