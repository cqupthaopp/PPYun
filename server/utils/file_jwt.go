package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//jwt Struct
type FileClaims struct {
	Path      string `json:"path"`
	FileOwner string `json:"owner-name"`
	TokenType string `json:"token-type"`
	jwt.StandardClaims
}

const (
	FileTokenTime = time.Hour * 1
	TokenType     = "FileToken"
) //过期时间和密钥

//GetToken get user's token
func CreateFileToken(username string, path string) (string, error) {
	c := FileClaims{path, username, TokenType, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(FileTokenTime).Unix(),
		Issuer:    issuer,
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(secret))
}

func PraseFileToken(tokenStr string) (*FileClaims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &FileClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*FileClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("PraseToken Error")
}

//JudgeAccessToken Return owner,path
func JudgeFileToken(msg string) (string, string) {

	claims, err := PraseFileToken(msg)

	if err != nil {
		return "NULL", "NULL"
	}
	if claims.StandardClaims.ExpiresAt <= time.Now().Unix() || claims.StandardClaims.Issuer != issuer || claims.TokenType != TokenType {
		return "NULL", "NULL"
	}

	return claims.FileOwner, claims.Path
}
