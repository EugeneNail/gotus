package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// A ResponseWriter is the interface that wraps the basic Write method
//
// Write writes data given from the response to the buffer.
// Write must return a non-nil error if there is a problem during data writing, otherwise nil
type ResponseWriter interface {
	Write(*Response, *bytes.Buffer, http.ResponseWriter) error
}

// A JsonResponseWriter writes data to a buffer as JSON.
type JsonResponseWriter struct{}

// Write sets the Content-Type header to "application/json", formats data and writes it to the buffer using a json.JsonEncoder
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

// A PlainResponseWriter writes data to a buffer as plain text
type PlainResponseWriter struct{}

// Write sets the Content-Type header to "text/plain" and writes plain text from the response to the buffer
func (PlainResponseWriter) Write(response *Response, buffer *bytes.Buffer, writer http.ResponseWriter) error {
	writer.Header().Set("Content-Type", "text/plain")

	if _, err := buffer.Write([]byte(response.Data.(string))); err != nil {
		return fmt.Errorf("cannot write text to buffer: %w", err)
	}

	return nil
}
