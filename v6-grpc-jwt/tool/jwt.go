package tool

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JwtSecret = []byte("demo111")

type Token struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

// 创建jwt token
func CreateJwtToken(Id int32, name string) (string, error) {
	var token Token
	token.StandardClaims = jwt.StandardClaims{
		Audience:  "",
		ExpiresAt: time.Now().Add(30 * time.Second).Unix(),
		Id:        "",
		IssuedAt:  time.Now().Unix(),
		Issuer:    "timeless",
		NotBefore: 0,
		Subject:   "login",
	}
	token.Name = name
	token.Id = Id
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, token)
	fmt.Println(tokenClaims)
	fmt.Println(JwtSecret)
	return tokenClaims.SignedString(JwtSecret)
}

// 解析 token 以及相关信息
func ParseJwtToken(token string) (Token, error) {
	var clientClaims Token
	jwtToken, err := jwt.ParseWithClaims(token, &clientClaims, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if jwtToken == nil || err != nil {
		return Token{}, err
	}
	if jwtToken.Valid {
		return clientClaims, nil
	}
	return Token{}, nil
}

type AuthToken struct {
	Token string
}

func (a AuthToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": a.Token,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires
// transport security.
func (a AuthToken) RequireTransportSecurity() bool {
	return false
}
