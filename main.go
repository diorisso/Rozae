package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"rozae/database"
	"rozae/handlers"
	"rozae/middlewares"
	"rozae/repositories"
	"rozae/services"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	const porta = "8080"
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao inicializar .env", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	connectionString := os.Getenv("DATABASE_URL")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*500)
	defer cancel()

	db, err := database.Conectar(ctx, connectionString)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados: ", err)
	}
	defer db.Close()

	// repositories
	pessoaRepository := repositories.NewPessoaRepository(db)

	// services
	jwtService := services.NewJwtService(jwtSecret)
	authService := services.NewAuthService(pessoaRepository, jwtService)

	// handlers
	mux := http.NewServeMux()
	api := http.NewServeMux()

	authHandler := handlers.NewAuthHandler(authService)
	authHandler.AdicionarRotas(api)

	mux.Handle("/api/", http.StripPrefix("/api", middlewares.Autenticar(jwtService, api)))

	fmt.Println("Ouvindo na porta " + porta)
	http.ListenAndServe(":"+porta, mux)
}
