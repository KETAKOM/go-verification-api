package handler

import (
	"net/http"
)

type helloHandler struct{}

func NewHelloHandler() *helloHandler {
	return &helloHandler{}
}

type Test struct {
	Message string "json: message"
}

func (h *helloHandler) HelloHandler(w http.ResponseWriter, r *http.Request) {

	test := Test{
		Message: "hello golang ",
	}

	ResponseJson(w, http.StatusOK, test)
	return
}
