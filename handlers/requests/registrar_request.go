package requests

type RegistrarRequest struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}
