package main

import (
	"fmt"
	"net/http"
)

type Http interface {
	Server() string
}

type Service struct{}

func (s Service) Server() string {
	return "Hello world"
}

type HttpHandler struct {
	service Service
}

func HelloHandler(service Service) *HttpHandler {
	return &HttpHandler{service: service}
}

func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := h.service.Server()
	fmt.Fprintf(w, message)
}

func main() {
	service := Service{}

	helloHandler := HelloHandler(service)
	http.Handle("/hello", helloHandler)

	fmt.Println("Server running on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
