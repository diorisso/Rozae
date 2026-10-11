package utils

import "errors"

var (
	ErrCredenciaisInvalidas = errors.New("Credenciais inválidas")
	ErrEmailJaCadastrado    = errors.New("O email informado já está cadastrado")
)
