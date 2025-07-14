package transport

import "net/http"

// A Request wraps the http.Request struct to be able to store a typed payload
type Request[T Payload] struct {
	Payload T
	*http.Request
}
