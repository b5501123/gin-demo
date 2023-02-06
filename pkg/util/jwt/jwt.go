package jwtUtil

import (
	"gin-demo/model/bo"
	"gin-demo/model/claims"
	"github.com/golang-jwt/jwt"
	"time"
)

func GenToken(user bo.UserBo, secretKey string) (string, error) {

	c := claims.UserClaims{
		user.Id,
		user.Account,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	// Choose specific algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// Choose specific Signature
	return token.SignedString(secretKey)
}
