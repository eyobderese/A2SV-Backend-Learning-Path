package infrastructure

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("your_jwt_secret")

func TotokenParser(authPartToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(authPartToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSecret, nil
	})

	if err != nil || !token.Valid {

		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {

		return nil, err
	}
	return claims, nil
}

func TokenGeneretor(claims map[string]interface{}) (string, error) {
	jwtMapClaims := jwt.MapClaims(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtMapClaims)

	jwtToken, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	// return the token
	return jwtToken, nil
}
