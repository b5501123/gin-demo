package claims

import (
	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	Id      int64  `json:"id"`
	Account string `json:"account"`
	jwt.StandardClaims
}
