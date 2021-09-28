package main

import (
	"github.com/kukwuka/Gos/develop/3_11/pkg/handlers"
	"github.com/kukwuka/Gos/develop/3_11/pkg/services"
	"net/http"
)

func main() {
	service := services.NewService()
	handler := handlers.NewHandler(service)

	mux := handler.InitRoutes()

	http.ListenAndServe(":8080", mux)
}
