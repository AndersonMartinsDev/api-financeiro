package usuario

type UsuarioDTO struct {
	ID     uint   `json:"id,omitempty"`
	Nome   string `json:"nome,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}
