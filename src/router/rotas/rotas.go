package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI          string
	Metodo       string
	Func         func(w http.ResponseWriter, r *http.Request)
	Autenticacao bool
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := despesasRotas
	rotas = append(rotas, envelopeRotas...)

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Func).Methods(rota.Metodo)
	}
	return r
}
