package main

import (
	"11/pkg/handlers"
	"11/pkg/services"
	"net/http"
)

func main() {
	service := services.NewService()
	handler := handlers.NewHandler(service)

	mux := handler.InitRoutes()

	http.ListenAndServe(":8080", mux)
}
