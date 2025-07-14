package middleware

import (
	"context"
	"github.com/EugeneNail/gotus/internal/transport"
	"github.com/google/uuid"
	"net/http"
)

// SetRequestId is a middleware that inject new UUID into context of the request
func SetRequestId[T transport.Payload](next transport.HandlerFunc[T]) transport.HandlerFunc[T] {
	return func(writer http.ResponseWriter, r *transport.Request[T]) *transport.Response {
		request := transport.Request[T]{
			Request: r.WithContext(context.WithValue(r.Context(), "id", uuid.New().String())),
		}

		return next(writer, &request)
	}
}
