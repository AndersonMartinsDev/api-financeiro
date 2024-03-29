package autenticacao

import (
	"api/src/commons/config"
	"api/src/models/usuario"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken token com as permissões de usuário
func CriarToken(usuario usuario.UsuarioLoginDto) (string, error) {
	//1234
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["username"] = usuario.Username
	permissoes["username_id"] = uint(usuario.ID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	token.Valid = true
	return token.SignedString([]byte(config.SecretKey)) //secret
}

// ValidarToken verifica se o token na requisição é valido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)

	if erro != nil {
		log.Printf("ERRO DE TOKEN %s", erro)
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return fmt.Errorf("token inválido")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}
	return []byte(config.SecretKey), nil
}

// ExtrairUsername
func ExtrairUsername(r *http.Request) (uint, error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)

	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := permissoes["username_id"].(float64)
		if erro != nil {
			return 0, erro
		}
		return uint(id), erro
	}
	return 0, nil
}
