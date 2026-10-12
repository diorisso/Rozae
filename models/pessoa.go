package models

import "time"

type Pessoa struct {
	Id              int
	Nome            string
	Email           string
	Senha           string
	EmailConfirmado bool
	Excluido        bool
	CriadoEm        time.Time
	ModificadoEm    time.Time
}
