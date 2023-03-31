package usuario

type UsuarioResumeDTO struct {
	ID       uint   `json:"id"`
	Nome     string `json:"nome"`
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
}
