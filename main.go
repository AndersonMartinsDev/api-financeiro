package main

import (
	"api/src/router"
	"api/src/tools/config"
	"fmt"
	"net/http"
)

func main() {
	config.Carregar()
	r := router.Gerar()
	fmt.Println("init aplicação ")
	http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r)
}
