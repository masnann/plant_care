package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"time"
)

type JWTInterface interface {
	GenerateJWT(userID uint64) (string, string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	GenerateRefreshJWT(accessToken *jwt.Token, refreshToken *jwt.Token) map[string]any
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
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

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
func (j *JWT) GenerateRefreshJWT(accessToken *jwt.Token, refreshToken *jwt.Token) map[string]any {
	var result = map[string]any{}
	expTime, err := refreshToken.Claims.GetExpirationTime()
	logrus.Info(expTime)
	if err != nil {
		logrus.Error("get token expiration error", err.Error())
		return nil
	}
	if refreshToken.Valid && expTime.Time.Compare(time.Now()) > 0 {
		var newClaim = accessToken.Claims.(jwt.MapClaims)

		newClaim["iat"] = time.Now().Unix()
		newClaim["exp"] = time.Now().Add(time.Hour * 1).Unix()

		var newToken = jwt.NewWithClaims(accessToken.Method, newClaim)
		newSignedToken, _ := newToken.SignedString(accessToken.Signature)

		var newRefreshClaim = refreshToken.Claims.(jwt.MapClaims)
		newRefreshClaim["exp"] = time.Now().Add(time.Hour * 24).Unix()

		var newRefreshToken = jwt.NewWithClaims(refreshToken.Method, newRefreshClaim)
		newSignedRefreshToken, _ := newRefreshToken.SignedString(refreshToken.Signature)

		result["access_token"] = newSignedToken
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
