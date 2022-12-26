package http

import "request/http/method"

type Request struct {
	Method  method.Method     `json:"method"`
	Url     string            `json:"url"`
	Params  map[string]string `json:"params"`
	Headers *[]Header
	Body    Body
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Body interface {
	Data() []byte
}
