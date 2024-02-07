package auth

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = os.Getenv("JWT_SECRET")

func CreateToken(email, userID string) (string, error) {
	claims := jwt.MapClaims{
		"email":  email,
		"userId": userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func VerifyToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid token")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("Invalid token")
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); !ok {
		return nil, fmt.Errorf("Invalid token")
	} else {
		return claims, nil
	}

}
