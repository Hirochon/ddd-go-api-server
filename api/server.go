package main

import (
	"api/handler"
	"api/infra"
	"api/middleware"
	"api/usecase"
	"fmt"
	"net/http"
)

var addr = ":8080"

func main() {
	myClient := infra.Init()
	defer myClient.Close()
	fmt.Println(myClient)

	statusRepository := infra.NewStatusRepository(myClient)
	statusUsecase := usecase.NewStatusUsecase(statusRepository)
	taskHandler := handler.NewStatusHandler(statusUsecase)

	router := http.NewServeMux()
	router.HandleFunc("/", taskHandler.POST())
	fmt.Printf("[START] server. port: %s\n", addr)
	if err := http.ListenAndServe(addr, middleware.Log(router)); err != nil {
		panic(fmt.Errorf("[FAILED] start sever. err: %v", err))
	}
}
