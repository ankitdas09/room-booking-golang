package main

import (
	"net/http"

	"github.com/ankitdas09/gowebapp/cmd/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	http.ListenAndServe(":8080", nil)
}
