package usuario

type UsuarioLoginDto struct {
	Username   string `json:"username"`
	Senha      string `json:"senha"`
	CarteiraId string `json:"carteiraId,omitempty"`
}
