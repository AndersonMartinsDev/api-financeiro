package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//SecretKey chave secreta
	SecretKey = "VS_3zaqZyL8pgJSiKqINkNTBcRn88Br7wA2eOLxTwQo"
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
	host := os.Getenv("DB_HOST")
	porta := os.Getenv("DB_PORT")
	db_config := "charset=utf8&parseTime=True&loc=Local"

	// Uri = fmt.Sprintf("%s:%s@/%s?%s", usuario, senha, nome, db_config)
	Uri = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", usuario, senha, host, porta, nome, db_config)
	fmt.Println(Uri)

	// fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
}
