package main

import (
	"github.com/EugeneNail/gotus/internal/service/log"
	"net"
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
	http.HandleFunc("/ping", PingPongHandler)
	http.Handle("/", newUiHandler())

	go log.Initialize()

	http.ListenAndServe(net.JoinHostPort("", os.Getenv("APP_PORT")), nil)
}
