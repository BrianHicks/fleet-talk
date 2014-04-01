package main

import (
	"fmt"
	"net/http"
	"os"
)

var host string

func init() {
	host = os.Getenv("HOST")
	if host == "" {
		host = "{HOST not set}"
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %s!", host)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
