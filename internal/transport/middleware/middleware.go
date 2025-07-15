package middleware

import (
	"github.com/EugeneNail/gotus/internal/transport"
	"net/http"
)

// A Middleware is a function that uses a transport.HandleFunc to process it and produce new transport.HandleFunc.
// Middlewares must be chained to create request pipelines.
type Middleware[T transport.Payload] func(next transport.HandlerFunc[T]) transport.HandlerFunc[T]

// Web provides a request pipeline with predefined set of middlewares often used in [browser-->server] requests
func Web[T transport.Payload](handler transport.HandlerFunc[T]) http.HandlerFunc {
	middlewares := []Middleware[T]{
		SetRequestId[T],
		WriteResponse[T],
		Validate[T],
	}

	return Chain(handler, middlewares)
}

// Chain creates a request pipeline from multiple middlewares and converts it to a type that can be used in Go's default routing API
func Chain[T transport.Payload](next transport.HandlerFunc[T], middlewares []Middleware[T]) http.HandlerFunc {
	length := len(middlewares)
	chained := middlewares[length-1](next)

	for i := length - 2; i >= 0; i-- {
		chained = middlewares[i](chained)
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		customRequest := &transport.Request[T]{Request: request}
		chained(writer, customRequest)
	}
}
