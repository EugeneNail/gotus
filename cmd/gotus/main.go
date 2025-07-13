package main

import (
	"github.com/EugeneNail/gotus/internal/service/log"
	"net"
	"net/http"
	"os"
	"path"
)

func main() {
	log.Initialize()

	router := http.NewServeMux()
	router.Handle(`/`, http.FileServer(http.Dir(path.Join(os.Getenv("APP_ROOT"), "web"))))

	log.Info("Server started")
	if err := http.ListenAndServe(net.JoinHostPort("", os.Getenv("APP_PORT")), router); err != nil {
		log.Error(err)
	}
}
