package routing

import (
	"github.com/EugeneNail/gotus/api"
	"net/http"
)

type HandlerFunc[T api.Payload] func(http.ResponseWriter, *Request[T])

func (f HandlerFunc[T]) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, &Request[T]{
		Request:   r,
		Variables: make(map[string]string),
	})
}
