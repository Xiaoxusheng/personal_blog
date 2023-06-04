package test

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"testing"
	"time"
)

type MyCustomClaims struct {
	Identification string `json:"identification"`
	jwt.RegisteredClaims
}

func Test_token(T *testing.T) {
	claims := MyCustomClaims{
		"123",
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	MySigningKey := []byte("jij i fe ")
	ss, _ := token.SignedString(MySigningKey)
	fmt.Println(ss)
}
