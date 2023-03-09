package autenticacao

import (
	"api/src/commons/config"
	"api/src/models/usuario"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken token com as permissões de usuário
func CriarToken(usuario usuario.UsuarioLoginDto) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["username"] = usuario.Username
	permissoes["carteiraId"] = usuario.CarteiraId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.Signature, nil //secret
}

// ValidarToken verifica se o token na requisição é valido
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)

	if erro != nil {
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
	return config.SecretKey, nil
}

// ExtrairUsuarioID
func ExtrairCarteiraId(r *http.Request) (string, error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)

	if erro != nil {
		return "", erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		carteira := fmt.Sprintf("%s", permissoes["carteira"])
		if erro != nil {
			return "", erro
		}

		return carteira, nil
	}
	return "", nil
}
