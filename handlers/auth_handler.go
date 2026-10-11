package handlers

import (
	"context"
	"errors"
	"net/http"
	"rozae/handlers/requests"
	"rozae/handlers/responses"
	"rozae/utils"
)

type authService interface {
	Login(ctx context.Context, email string, senha string) (*responses.RegistrarResponse, error)
	Registrar(ctx context.Context, email string, senha string) (*responses.RegistrarResponse, error)
}

type AuthHandler struct {
	service authService
}

func (h *AuthHandler) AdicionarIoC(mux *http.ServeMux) {
	mux.HandleFunc("POST /auth/login", utils.WithJson(h.Login))
	mux.HandleFunc("POST /auth/registrar", utils.WithJson(h.Login))
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request, obj requests.LoginRequest) {
	retorno, err := h.service.Login(r.Context(), obj.Email, obj.Senha)

	if err != nil {
		if errors.Is(err, utils.ErrCredenciaisInvalidas) {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, http.StatusOK, retorno)
}

func (h *AuthHandler) Registrar(w http.ResponseWriter, r *http.Request, obj requests.RegistrarRequest) {
	retorno, err := h.service.Registrar(r.Context(), obj.Email, obj.Senha)

	if err != nil {
		if errors.Is(err, utils.ErrEmailJaCadastrado) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, http.StatusCreated, retorno)
}
