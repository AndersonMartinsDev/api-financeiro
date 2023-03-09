package main

import (
	"api/src/commons/config"
	"api/src/router"
	"fmt"
	"net/http"
)

func main() {
	config.Carregar()
	r := router.Gerar()
	fmt.Println("init aplicação ")
	http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r)
}
