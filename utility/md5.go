package utility

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
	"personal_blog/db"
	"time"
)

// SetUuid 生成uuid
func SetUuid() string {
	fmt.Println(uuid.NewV4().String())
	return uuid.NewV4().String()
}

// GetMa5 生成MD5
func GetMa5(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

type MyCustomClaims struct {
	Identification string `json:"identification"`
	jwt.RegisteredClaims
}

// GetToken 生成token
func GetToken(identification string) string {
	ctx := context.Background()
	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		identification,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(MySigningKey)
	//生成token同时存入redis
	result, err := db.Rdb.Set(ctx, identification, ss, time.Hour*24).Result()
	if err != nil {
		return ""
	}
	fmt.Println(result)

	fmt.Printf("%v ", ss)
	return ss
}
