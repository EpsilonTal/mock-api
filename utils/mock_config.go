package utils

type MockConfig struct {
	ID     string       `json:"id"`
	Post   MethodConfig `json:"post"`
	Delete MethodConfig `json:"delete"`
}

type MethodConfig struct {
	Status int64       `json:"status"`
	Body   interface{} `json:"body"`
}
