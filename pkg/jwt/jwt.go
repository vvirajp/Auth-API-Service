package jwt

import (
	"fmt"
	"math/rand"

	"github.com/golang-jwt/jwt/v5"
)

type JWTTokenData struct {
	User    string `json:"user"`
	ExpTime int64
}

func GenerateToken(tokenData JWTTokenData) (string, int64, error) {
	secretKey := []byte("ti1iu1aal81sx84fjkd4aiirq0zth7ol")
	randomNum := rand.Intn(900000000000) + 112345654321

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":      tokenData.User,
		"exp":       tokenData.ExpTime,
		"randomNum": randomNum,
	})

	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", 0, err
	}

	return accessToken, tokenData.ExpTime, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	secretKey := []byte("ti1iu1aal81sx84fjkd4aiirq0zth7ol")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("couldn't parse token claims")
	}

	return claims, nil
}
