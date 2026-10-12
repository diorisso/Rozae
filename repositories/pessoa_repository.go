package repositories

import (
	"context"
	"database/sql"
	"errors"
	"rozae/models"
	"time"
)

type PessoaRepository struct {
	db *sql.DB
}

func NewPessoaRepository(db *sql.DB) *PessoaRepository {
	return &PessoaRepository{
		db: db,
	}
}

func (r *PessoaRepository) ObterPessoaPorEmail(ctx context.Context, email string) (*models.Pessoa, error) {
	var pessoa models.Pessoa
	query := `SELECT id, email, senha FROM pessoas WHERE email LIKE $1 AND excluido = false`

	err := r.db.QueryRowContext(
		ctx,
		query,
		email,
	).Scan(
		&pessoa.Id,
		&pessoa.Email,
		&pessoa.Senha,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &pessoa, nil
}

func (r *PessoaRepository) CriarNovaPessoa(ctx context.Context, email string, nome string, senha string) (int, error) {
	pessoa := models.Pessoa{
		Nome:            nome,
		Email:           email,
		Senha:           senha,
		EmailConfirmado: false,
		Excluido:        false,
		CriadoEm:        time.Now().UTC(),
		ModificadoEm:    time.Now().UTC(),
	}

	query := `
		INSERT INTO pessoas (nome, email, senha, email_confirmado, excluido, criado_em, modificado_em)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	var id int
	err := r.db.QueryRowContext(
		ctx,
		query,
		pessoa.Nome,
		pessoa.Email,
		pessoa.Senha,
		pessoa.EmailConfirmado,
		pessoa.Excluido,
		pessoa.CriadoEm,
		pessoa.ModificadoEm,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
