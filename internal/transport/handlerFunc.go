package transport

import "net/http"

// A HandlerFunc wraps the http.HandleFunc function into a function that is able to pass responses to underlying middlewares.
type HandlerFunc[T Payload] func(writer http.ResponseWriter, request *Request[T]) *Response
