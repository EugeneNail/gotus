package routing

import (
	"github.com/EugeneNail/gotus/api"
	"github.com/EugeneNail/gotus/internal/enum/method"
	"net/http"
	"regexp"
)

type Route struct {
	method  method.Method
	pattern *regexp.Regexp
	handler http.Handler
	payload api.Payload
}
