package main

import (
	"fmt"
	"net/http"
	"rozae/middlewares"
	"rozae/services"
)

func main() {
	const porta = "8080"

	mux := http.NewServeMux()
	api := http.NewServeMux()

	jwtService := services.NewJwtService("secretsupersecretoblablabla")

	mux.Handle("/api/", http.StripPrefix("/api", middlewares.Autenticar(jwtService, api)))

	fmt.Println("Ouvindo na porta " + porta)
	http.ListenAndServe(":"+porta, mux)
}
