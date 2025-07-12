package routing

import (
	"github.com/EugeneNail/gotus/api"
	"github.com/EugeneNail/gotus/internal/core/log"
	"github.com/EugeneNail/gotus/internal/enum/method"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var routes = make([]*Route, 0)

func Get[T api.Payload](url string, handler HandlerFunc[T]) {
	AddRoute[T](method.Get, url, handler)
}

func Post[T api.Payload](url string, handler HandlerFunc[T]) {
	AddRoute[T](method.Post, url, handler)
}

func Put[T api.Payload](url string, handler HandlerFunc[T]) {
	AddRoute[T](method.Put, url, handler)
}

func Patch[T api.Payload](url string, handler HandlerFunc[T]) {
	AddRoute[T](method.Patch, url, handler)
}

func Delete[T api.Payload](url string, handler HandlerFunc[T]) {
	AddRoute[T](method.Delete, url, handler)
}

func AddRoute[T api.Payload](method method.Method, url string, handler HandlerFunc[T]) {
	pattern := regexp.MustCompile("{[0-9a-zA-Z]+}").ReplaceAllString(url, "[0-9a-zA-Z]+")
	var payload T

	routes = append(routes, &Route{
		method:  method,
		handler: handler,
		pattern: regexp.MustCompile("^" + pattern + "$"),
		payload: payload,
	})
}

func Start() {
	http.HandleFunc("/api/", func(writer http.ResponseWriter, request *http.Request) {
		defer log.RedirectPanicToLogger()
		allowedMethods := make([]string, 0)

		for _, route := range routes {
			if !route.pattern.MatchString(request.URL.String()) {
				continue
			}

			if route.method.ToString() == strings.ToUpper(request.Method) {
				route.handler.ServeHTTP(writer, request)
				return
			}

			allowedMethods = append(allowedMethods, route.method.ToString())
		}

		if len(allowedMethods) == 0 {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		writer.Header().Set("Allow", strings.Join(allowedMethods, ", "))
		writer.WriteHeader(http.StatusMethodNotAllowed)
	})

	log.Info("Server started")
	err := http.ListenAndServe(net.JoinHostPort("", os.Getenv("APP_PORT")), nil)
	if err != nil {
		panic(err)
	}
}
