package main

import (
	"fmt"
	"net/http"
)

func Dock(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Docker container!<h1>")
}

func main() {
	http.HandleFunc("/", Dock)
	http.ListenAndServe(":1111", nil)
}
