package usuario

type UsuarioLoginDto struct {
	ID       uint
	Username string `json:"username"`
	Senha    string `json:"senha"`
}
