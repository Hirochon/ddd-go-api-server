package main

import (
	"api/infra"
	"api/middleware"
	"fmt"
	"net/http"
)

var addr = ":8080"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	myClient := infra.Init()
	defer myClient.Close()
	fmt.Println(myClient)

	router := http.NewServeMux()
	router.HandleFunc("/", handler)
	fmt.Printf("[START] server. port: %s\n", addr)
	if err := http.ListenAndServe(addr, middleware.Log(router)); err != nil {
		panic(fmt.Errorf("[FAILED] start sever. err: %v", err))
	}
}
