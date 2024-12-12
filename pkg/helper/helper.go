package helper

import (
	"API/pkg/config"
	"API/pkg/jwt"
	"fmt"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ValidateEmail(email string) bool {
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailRegex := regexp.MustCompile(emailRegexPattern)
	return emailRegex.MatchString(email)
}

func Authenticate(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")

	// Check if Authorization header is present
	if authHeader == "" {
		return "", fmt.Errorf("missing authorization header")
	}

	jwtClaims, err := jwt.VerifyToken(authHeader)
	if err != nil {
		return "", fmt.Errorf("invalid token")
	}
	emailJWT := jwtClaims["user"].(string)

	cache := config.GetCache()
	val, found := cache.Get(authHeader)
	if !found {
		return "", fmt.Errorf("Invalid_Token")
	} else if val == emailJWT {
		return emailJWT, nil
	} else {
		return "", fmt.Errorf("token authorization failed")
	}

}

func AuthenticateRefresh(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")

	// Check if Authorization header is present
	if authHeader == "" {
		return "", fmt.Errorf("missing authorization header")
	}

	jwtClaims, err := jwt.VerifyToken(authHeader)
	if err != nil {
		return "", fmt.Errorf("invalid token")
	}
	emailJWT := jwtClaims["user"].(string)

	cache := config.GetCache()
	val, found := cache.Get(authHeader)
	if !found {
		return "", fmt.Errorf("Invalid_Token")
	} else if val == emailJWT && strings.Contains(val.(string), "refreshtoken") {
		return emailJWT, nil
	} else if val == emailJWT && !strings.Contains(val.(string), "refreshtoken") {
		return "", fmt.Errorf("1")
	} else {
		return "", fmt.Errorf("token authorization failed")
	}

}
