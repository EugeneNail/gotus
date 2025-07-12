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

func main() {
	log.Initialize()
	http.Handle("/", newUiHandler())

	routing.Start()
}
