package main

import (
	"net"
	"net/http"
	"os"
)

func newUiHandler() http.Handler {
	return http.FileServer(http.Dir(os.Getenv("APP_ROOT") + "/web"))
}

func PingPongHandler(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte("pong")); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/ping", PingPongHandler)

	http.Handle("/", newUiHandler())

	http.ListenAndServe(net.JoinHostPort("", os.Getenv("APP_PORT")), nil)
}
