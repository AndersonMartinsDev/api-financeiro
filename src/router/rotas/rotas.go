package rotas

import (
	"api/src/commons/middlewares"
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
	rotas := authRoutes
	rotas = append(rotas, homeRotas...)
	rotas = append(rotas, despesasRotas...)
	rotas = append(rotas, usuarioRotas...)
	rotas = append(rotas, envelopeRotas...)
	rotas = append(rotas, pagamentosRotas...)

	for _, rota := range rotas {
		if rota.Autenticacao {
			r.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Func))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.Func)).Methods(rota.Metodo)
		}
	}
	return r
}
