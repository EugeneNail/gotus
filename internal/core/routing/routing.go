package routing

import (
	"github.com/EugeneNail/gotus/internal/core/log"
	"github.com/EugeneNail/gotus/internal/enum/method"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var routes = make([]*Route, 0)

func Start() {
	http.HandleFunc("/api/", func(writer http.ResponseWriter, request *http.Request) {
		defer log.RedirectPanicToLogger()
		allowedMethods := make([]string, 0)

		for _, route := range routes {
			if route.pattern.MatchString(request.URL.String()) {
				if route.method.ToString() == strings.ToUpper(request.Method) {
					route.handler.ServeHTTP(writer, request)
					return
				}

				allowedMethods = append(allowedMethods, route.method.ToString())
			}
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

func Get(url string, handler http.HandlerFunc) {
	AddRoute(method.Get, url, handler)
}

func Post(url string, handler http.HandlerFunc) {
	AddRoute(method.Post, url, handler)
}

func Put(url string, handler http.HandlerFunc) {
	AddRoute(method.Put, url, handler)
}

func Patch(url string, handler http.HandlerFunc) {
	AddRoute(method.Patch, url, handler)
}

func Delete(url string, handler http.HandlerFunc) {
	AddRoute(method.Delete, url, handler)
}

func AddRoute(method method.Method, url string, handler http.HandlerFunc) {
	pattern := regexp.MustCompile("{[0-9a-zA-Z]+}").ReplaceAllString(url, "[0-9a-zA-Z]+")
	routes = append(routes, &Route{
		method:  method,
		handler: handler,
		pattern: regexp.MustCompile("^" + pattern + "$"),
	})
}
