package main

import (
	"github.com/EugeneNail/gotus/internal/core/log"
	"github.com/EugeneNail/gotus/internal/core/routing"
	"net/http"
	"os"
	"path"
)

func newUiHandler() http.Handler {
	return http.FileServer(http.Dir(path.Join(os.Getenv("APP_ROOT"), "web")))
}

func PingPongHandler(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte("pong")); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	log.Initialize()
	http.Handle("/", newUiHandler())

	routing.Start()
}

func response(message string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(message))
	}
}
