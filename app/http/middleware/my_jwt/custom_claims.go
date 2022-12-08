package my_jwt

import (
	"github.com/dgrijalva/jwt-go"
)

type AccClaims struct {
	Id      int64  `json:"id"`
	Account string `json:"account"`
	Status  int64  `json:"status"`
	jwt.StandardClaims
}

func (accClaims *AccClaims) IsAccClaimsValid() bool {
	if accClaims.Id == 0 {
		return false
	}
	return true
}
