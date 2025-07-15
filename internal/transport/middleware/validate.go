package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/EugeneNail/gotus/internal/service/validation"
	"github.com/EugeneNail/gotus/internal/transport"
	"net/http"
)

func Validate[T transport.Payload](next transport.HandlerFunc[T]) transport.HandlerFunc[T] {
	return func(writer http.ResponseWriter, request *transport.Request[T]) *transport.Response {
		var data map[string]any
		if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
			return transport.NewPlainResponse(http.StatusBadRequest, fmt.Sprintf("malformed json: %s", err.Error()))
		}

		var lastError validation.Error
		errors := make(map[string]string)
		for field, rules := range request.Payload.Rules() {
			for _, ruleFunc := range rules {
				if err := ruleFunc(field, data); err != nil {
					errors[field] = err.Error()
					lastError = err
					break
				}
			}
		}

		if len(errors) > 0 {
			return transport.NewJsonResponse(http.StatusUnprocessableEntity, lastError.Error(), errors)
		}

		return next(writer, request)
	}
}
