package services

import (
	"context"
	"rozae/handlers/responses"
	"rozae/models"
	"rozae/repositories"
	"rozae/utils"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type authRepository interface {
	ObterPessoaPorEmail(ctx context.Context, email string) (*models.Pessoa, error)
	CriarNovaPessoa(ctx context.Context, email string, nome string, senha string) (int, error)
}

type AuthService struct {
	repository authRepository
	jwtService *JwtService
}

func NewAuthService(repo *repositories.PessoaRepository, jwtService *JwtService) *AuthService {
	return &AuthService{
		repository: repo,
		jwtService: jwtService,
	}
}

func (s *AuthService) Login(ctx context.Context, email string, senha string) (*responses.LoginResponse, error) {
	pessoa, err := s.repository.ObterPessoaPorEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if pessoa == nil {
		return nil, utils.ErrCredenciaisInvalidas
	}

	err = bcrypt.CompareHashAndPassword([]byte(pessoa.Senha), []byte(senha))
	if err != nil {
		return nil, utils.ErrCredenciaisInvalidas
	}

	token, err := s.jwtService.GerarToken(strconv.Itoa(pessoa.Id))
	if err != nil {
		return nil, err
	}

	retorno := responses.LoginResponse{
		AccessToken: token,
		ExpiresIn:   3600, // TODO: botar nas configs
	}

	return &retorno, nil
}

func (s *AuthService) Registrar(ctx context.Context, email string, nome string, senha string) (*responses.RegistrarResponse, error) {
	pessoaJaExistente, err := s.repository.ObterPessoaPorEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	if pessoaJaExistente != nil {
		return nil, utils.ErrEmailJaCadastrado
	}

	senhaHashed, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	novaPessoaId, err := s.repository.CriarNovaPessoa(ctx, email, nome, string(senhaHashed))
	if err != nil {
		return nil, err
	}

	token, err := s.jwtService.GerarToken(strconv.Itoa(novaPessoaId))
	if err != nil {
		return nil, err
	}

	retorno := responses.RegistrarResponse{
		AccessToken: token,
		ExpiresIn:   3600,
	}

	return &retorno, nil
}
