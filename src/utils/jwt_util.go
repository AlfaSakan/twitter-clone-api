package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"github.com/AlfaSakan/twitter-clone-api/src/helpers"

	"github.com/golang-jwt/jwt/v4"
)

type CustomClaim struct {
	*entities.User
	SessionId string `json:"session_id"`
}

func GenerateToken(data *CustomClaim, expireTime int64) (string, error) {
	privateKey := os.Getenv("PRIVATEKEY")

	claims := jwt.MapClaims{
		"iss": "issuer",
		"exp": expireTime,
		"data": map[string]interface{}{
			"name":       data.Name,
			"id":         data.Id,
			"email":      data.Email,
			"username":   data.Username,
			"created_at": data.CreatedAt,
			"updated_at": data.UpdatedAt,
			"session_id": data.SessionId,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString([]byte(privateKey))

	if err != nil {
		return "", err
	}

	return token, nil
}

func DecodeToken(token string) (jwt.MapClaims, error) {
	privateKey := os.Getenv("PRIVATEKEY")

	decode, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(privateKey), t.Claims.Valid()
	})

	if err != nil {
		return nil, err
	}

	claims := decode.Claims.(jwt.MapClaims)

	expiredMilliSecond := helpers.ConvertFloat64ToInt64(claims["exp"].(float64))

	isExpired := expiredMilliSecond-time.Now().UnixMilli() <= 0

	if isExpired {
		return claims, fmt.Errorf("token is expired")
	}

	return claims, err
}
