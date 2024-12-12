package auth

import (
	"API/pkg/config"
	"API/pkg/jwt"
	"fmt"
	"time"
)

func Signin(email, password string) (string, string, error) {
	storedPassword := cacheReader(email)
	if storedPassword == "" {
		return "", "", fmt.Errorf("1")
	}
	if storedPassword != password {
		return "", "", fmt.Errorf("1")
	}
	//generate new access token for this user
	token, _, err := jwt.GenerateToken(jwt.JWTTokenData{
		User:    email,
		ExpTime: time.Now().Add(5 * time.Minute).Unix(),
	})
	if err != nil {
		return "", "", err
	}

	//store token in memory
	err = cacheWriter(token, email)
	if err != nil {
		return "", "", err
	}

	//generate new refresh token for this user
	refreshToken, _, err := jwt.GenerateToken(jwt.JWTTokenData{
		User:    email + "refreshtoken",
		ExpTime: time.Now().Add(5 * time.Hour).Unix(),
	})
	if err != nil {
		return "", "", err
	}
	err = cacheWriter(refreshToken, email+"refreshtoken")
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func Signup(body SignupDTO) error {
	err := cacheWriter(body.Email, body.Password)
	if err != nil {
		return fmt.Errorf("1")
	}
	return nil
}

func RefreshToken(email string) (string, error) {
	//generate new refresh token for this user
	token, _, err := jwt.GenerateToken(jwt.JWTTokenData{
		User:    email,
		ExpTime: time.Now().Add(5 * time.Hour).Unix(),
	})
	if err != nil {
		return "", err
	}
	err = cacheWriter(token, email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func RevokeToken(token string) {
	c := config.GetCache()
	c.Delete(token)
}

func cacheWriter(key, value string) error {
	c := config.GetCache()
	err := c.Add(key, value, 0)
	if err != nil {
		return err
	}
	return nil
}

func cacheReader(key string) string {
	c := config.GetCache()
	val, found := c.Get(key)
	if found {
		return val.(string)
	} else {
		return ""
	}
}
