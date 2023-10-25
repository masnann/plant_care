package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/gommon/log"
	"github.com/masnann/plant_care/config"
	"github.com/sirupsen/logrus"
	"time"
)

type JWTInterface interface {
	GenerateJWT(userID uint64) (string, string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	GenerateRefreshJWT(accessToken string, refreshToken *jwt.Token, signKey string) map[string]any
}

type JWT struct {
	signKey    string
	refreshKey string
}

func NewJWT(signKey string, refreshKey string) JWTInterface {
	return &JWT{
		signKey:    signKey,
		refreshKey: refreshKey,
	}
}

func (j *JWT) GenerateJWT(userID uint64) (string, string, error) {
	var accessToken = j.GenerateToken(userID)
	if accessToken == "" {
		return "", "", errors.New("failed to generate access token")
	}

	refreshToken, err := j.generateRefreshToken(userID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (j *JWT) GenerateToken(id uint64) string {
	var claims = jwt.MapClaims{}
	claims["id"] = id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := sign.SignedString([]byte(j.signKey))

	if err != nil {
		return ""
	}

	return validToken
}

func (j *JWT) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.signKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
func (j *JWT) GenerateRefreshJWT(accessToken string, refreshToken *jwt.Token, signKey string) map[string]any {
	var result = map[string]any{}
	expTime, err := refreshToken.Claims.GetExpirationTime()
	logrus.Info(expTime)
	if err != nil {
		logrus.Error("get token expiration error", err.Error())
		return nil
	}
	if refreshToken.Valid && expTime.Time.Compare(time.Now()) > 0 {
		var newClaim = jwt.MapClaims{}

		newToken, err := jwt.ParseWithClaims(accessToken, newClaim, func(t *jwt.Token) (interface{}, error) {
			return []byte(signKey), nil
		})

		if err != nil {
			log.Error(err.Error())
			return nil
		}

		newClaim = newToken.Claims.(jwt.MapClaims)
		newClaim["iat"] = time.Now().Unix()
		newClaim["exp"] = time.Now().Add(time.Hour * 1).Unix()

		var newRefreshClaim = refreshToken.Claims.(jwt.MapClaims)
		newRefreshClaim["exp"] = time.Now().Add(time.Hour * 24).Unix()

		var newRefreshToken = jwt.NewWithClaims(refreshToken.Method, newRefreshClaim)
		newSignedRefreshToken, _ := newRefreshToken.SignedString(refreshToken.Signature)

		result["access_token"] = newToken.Raw
		result["refresh_token"] = newSignedRefreshToken
		return result
	}

	return nil
}

func (j *JWT) generateRefreshToken(userID uint64) (string, error) {
	var claims = jwt.MapClaims{}
	claims["id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := sign.SignedString([]byte(j.refreshKey))

	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func ExtractToken(accessToken string) map[string]any {
	cfg := config.InitConfig()

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(accessToken, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
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
