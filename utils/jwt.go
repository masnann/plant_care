package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/gommon/log"
	"github.com/masnann/plant_care/config"
	"time"
)

type JWTInterface interface {
	GenerateJWT(userID uint64) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	ExtractToken(accessToken string) map[string]any
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) JWTInterface {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) GenerateJWT(userID uint64) (string, error) {
	var accessToken = j.GenerateToken(userID)
	if accessToken == "" {
		return "", errors.New("failed to generate access token")
	}

	return accessToken, nil
}

func (j *JWT) GenerateToken(id uint64) string {
	var claims = jwt.MapClaims{}
	claims["user_id"] = id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := sign.SignedString([]byte(j.Secret))

	if err != nil {
		return ""
	}

	return validToken
}

func (j *JWT) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
func (j *JWT) ExtractToken(accessToken string) map[string]any {
	cfg := config.InitConfig()

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(accessToken, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(cfg.Secret), nil
	})

	if err != nil {
		log.Error("Error Parsing Token : ", err.Error())
		return nil
	}

	if token.Valid {
		return map[string]any{
			"user-id": claims["id"],
		}
	}

	return nil
}
