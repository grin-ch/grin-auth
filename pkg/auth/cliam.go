package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type RoleBase struct {
	Id       int    `json:"id"`
	UUID     string `json:"uuid"`
	Avatar   string `json:"avator"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Time     int64  `json:"time"`
	Ip       string `json:"ip"`
}

type Cliams struct {
	RoleBase
	jwt.StandardClaims
}

func GenerateJWT(expires int, signed, issuer string, rBase RoleBase) (string, error) {
	now := time.Now()
	expire := now.Add(time.Duration(expires) * time.Second)
	rBase.Time = now.Unix()
	claims := Cliams{
		RoleBase: rBase,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			Issuer:    issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(signed))
	return token, err
}

func ParseJWT(token string) (*Cliams, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Cliams{}, func(token *jwt.Token) (interface{}, error) {
		return []byte{}, nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Cliams); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
