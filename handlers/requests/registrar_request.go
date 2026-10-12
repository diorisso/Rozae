package requests

type RegistrarRequest struct {
	Email string `json:"email"`
	Nome  string `json:"nome"`
	Senha string `json:"senha"`
}
