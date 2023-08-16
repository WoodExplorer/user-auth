package pkg

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserInfo struct {
	Name      string
	RoleNames []string
}

type UserInfoClaims struct {
	*jwt.StandardClaims
	UserInfo
}

func CreateToken(userInfo UserInfo, signKey string) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("HS256"))
	t.Claims = &UserInfoClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // TODO: use config
		},
		userInfo,
	}
	return t.SignedString([]byte(signKey))
}

func ParseToken(tokenString string, signKey string) (res UserInfo, err error) {
	tokenString = strings.TrimSpace(tokenString)

	token, err := jwt.ParseWithClaims(tokenString, &UserInfoClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	})
	if err != nil {
		return
	}

	claims := token.Claims.(*UserInfoClaims)
	return claims.UserInfo, nil
}
