package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		if _, err := writer.Write([]byte("pong")); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println(":" + os.Getenv("APP_PORT"))

	http.ListenAndServe(":"+os.Getenv("APP_PORT"), nil)
}
