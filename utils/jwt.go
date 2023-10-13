package utils

import (
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(payload jwt.MapClaims) string {
	viper := getYmlConfig()
	jwtConfig := viper.GetStringMap("jwt")
	secret := jwtConfig["secret"].(string)

	key := []byte(secret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, _ := token.SignedString(key)

	return tokenString
}
