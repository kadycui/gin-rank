package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kadycui/gin-rank/model"
)

var jwtKey = []byte("a_secret_key")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 分发Token
func ReleaseToken(user model.User) (string, error) {
	// 设置过期时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "KadyCui",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	return token, claims, err

}
