package routing

import (
	"github.com/EugeneNail/gotus/api"
	"net/http"
)

type Request[T api.Payload] struct {
	*http.Request
	Variables map[string]string
	Payload   T
}
