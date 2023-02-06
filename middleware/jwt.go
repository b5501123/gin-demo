package middleware

import (
	"errors"
	"gin-demo/config"
	"gin-demo/model/claims"
	"gin-demo/model/res"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, res.Error(401, "Authorization Empty"))
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, res.Error(401, "Authorization Empty"))
			c.Abort()
			return
		}

		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "Invalid Token.",
			})
			c.Abort()
			return
		}
		c.Set("user", &mc)

		c.Next()
	}
}

func ParseToken(tokenString string) (*claims.UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims.UserClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return config.Setting.WebConfig.JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// Valid token
	if claims, ok := token.Claims.(*claims.UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
