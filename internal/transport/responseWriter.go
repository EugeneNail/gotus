package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseWriter interface {
	Write(*Response, *bytes.Buffer, http.ResponseWriter) error
}

type JsonResponseWriter struct{}

func (JsonResponseWriter) Write(response *Response, buffer *bytes.Buffer, writer http.ResponseWriter) error {
	writer.Header().Set("Content-Type", "application/json")

	formatted := map[string]any{
		"message": response.Message,
		"data":    response.Data,
	}

	if err := json.NewEncoder(buffer).Encode(formatted); err != nil {
		return fmt.Errorf("cannot encode json to buffer: %w", err)
	}

	return nil
}

type PlainResponseWriter struct{}

func (PlainResponseWriter) Write(response *Response, buffer *bytes.Buffer, writer http.ResponseWriter) error {
	writer.Header().Set("Content-Type", "text/plain")

	if _, err := buffer.Write([]byte(response.Data.(string))); err != nil {
		return fmt.Errorf("cannot write text to buffer: %w", err)
	}

	return nil
}
