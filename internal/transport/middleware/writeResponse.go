package middleware

import (
	"bytes"
	"fmt"
	"github.com/EugeneNail/gotus/internal/service/log"
	"github.com/EugeneNail/gotus/internal/transport"
	"net/http"
)

func WriteResponse(next func(*http.Request) *transport.Response) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		response := next(request)
		var buffer bytes.Buffer

		err := response.Writer.Write(response, &buffer, writer)
		if err != nil {
			handleError(writer, request, err)
			return
		}

		writer.WriteHeader(response.Status)
		if _, err := buffer.WriteTo(writer); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		if response.Status >= 400 {
			log.Error(buildErrorMessage(response.Message, request))
		}
	}
}

func handleError(writer http.ResponseWriter, request *http.Request, err error) {
	log.Error(buildErrorMessage(err.Error(), request))
	writer.WriteHeader(http.StatusInternalServerError)
	if _, err := writer.Write([]byte(err.Error())); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func buildErrorMessage(message string, request *http.Request) string {
	var userId int
	if value := request.Context().Value("userId"); value != nil {
		userId = value.(int)
	}

	if userId > 0 {
		return fmt.Sprintf(`[%5d] %s at %s`, userId, message, request.URL.String())
	}

	return fmt.Sprintf(`%s at %s`, message, request.URL.String())
}
