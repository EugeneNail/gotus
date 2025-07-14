package middleware

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

func SetRequestId(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		next(writer, request.WithContext(context.WithValue(request.Context(), "id", uuid.New().String())))
	}
}
