package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

/**
 * @Author: Hao_pp
 * @Data: 2022-7-30 10:05
 * @Desc: JWTUtil
 */

//jwt Struct
type MyClaims struct {
	Username  string `json:"username"`
	TokenType string `json:"token-type"`
	jwt.StandardClaims
}

const (
	tokentime = time.Hour * 5
	secret    = "PPYun Author:Hao_pp"
	issuer    = "PPYun"
) //过期时间和密钥

//GetToken get user's token
func GetToken(username string) (string, error) {
	c := MyClaims{
		username, "AccessToken", jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokentime).Unix(),
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(secret))
}

func PraseToken(tokenStr string) (*MyClaims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("PraseToken Error")
}

//JudgeAccessToken Return user's username by his token
func JudgeAccessToken(msg string) (string, bool) {

	claims, err := PraseToken(msg)

	if err != nil {
		return "NULL", false
	}
	if claims.StandardClaims.ExpiresAt <= time.Now().Unix() || claims.StandardClaims.Issuer != issuer || claims.TokenType != "AccessToken" {
		return "NULL", false
	}

	return claims.Username, true
}
