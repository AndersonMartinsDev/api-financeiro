package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//SecretKey chave secreta
	SecretKey []byte
	//Porta seta porta para rodar a aplicação
	Porta = 0
	//URI possui o caminho do banco de dados
	Uri = ""
)

func Carregar() {
	var erro error
	if erro := godotenv.Load(); erro != nil {
		fmt.Errorf("Erro ao carregar variáveis de ambiente")
		return
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))

	if erro != nil {
		Porta = 9000
	}
	usuario := os.Getenv("DB_USUARIO")
	senha := os.Getenv("DB_SENHA")
	nome := os.Getenv("DB_NOME")
	db_config := "charset=utf8&parseTime=True&loc=Local"

	Uri = fmt.Sprintf("%s:%s@/%s?%s", usuario, senha, nome, db_config)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
