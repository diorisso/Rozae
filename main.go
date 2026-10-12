package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rozae/middlewares"
	"rozae/services"

	"github.com/joho/godotenv"
)

func main() {
	const porta = "8080"
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao inicializar .env", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	mux := http.NewServeMux()
	api := http.NewServeMux()

	jwtService := services.NewJwtService(jwtSecret)

	mux.Handle("/api/", http.StripPrefix("/api", middlewares.Autenticar(jwtService, api)))

	fmt.Println("Ouvindo na porta " + porta)
	http.ListenAndServe(":"+porta, mux)
}
