package usuario

type Usuario struct {
	ID       uint   `json:"id"`
	Avatar   string `json:"avatar,omitempty"`
	Nome     string `json:"nome"`
	Username string `json:"username"`
	Senha    string
	Email    string `json:"email,omitempty"`
}
