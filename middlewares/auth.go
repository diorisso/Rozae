package middlewares

import (
	"context"
	"net/http"
	"rozae/services"
	"strings"
)

type contextKey string

const PessoaIDKey contextKey = "pessoaID"

func Autenticar(
	jwtService *services.JwtService,
	next http.Handler,
) http.Handler {
	rotasPublicas := map[string]bool{
		"POST /auth/login":     true,
		"POST /auth/registrar": true,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rota := r.Method + " " + r.URL.Path

		if rotasPublicas[rota] {
			next.ServeHTTP(w, r)
			return
		}

		header := r.Header.Get("Authorization")
		partes := strings.Fields(header)

		if len(partes) != 2 ||
			!strings.EqualFold(partes[0], "Bearer") {
			http.Error(w, "Não autorizado", http.StatusUnauthorized)
			return
		}

		claims, err := jwtService.ValidarToken(partes[1])
		if err != nil {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			PessoaIDKey,
			claims.Subject,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
