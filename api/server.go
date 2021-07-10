package main

import (
	"api/infra"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	infra.Init()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
