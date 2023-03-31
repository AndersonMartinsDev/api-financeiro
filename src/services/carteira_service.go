package services

import (
	"golang.org/x/crypto/bcrypt"
)

func NovaHashCarteira(carteiraId string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(carteiraId), bcrypt.MinCost)

}
