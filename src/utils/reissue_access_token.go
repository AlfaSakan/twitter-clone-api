package utils

import (
	"time"

	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"github.com/AlfaSakan/twitter-clone-api/src/repositories"
	"gorm.io/gorm"
)

func ReIssueAccessToken(db *gorm.DB, refreshToken string) (string, *entities.User) {
	sessionRepository := repositories.NewSessionRepository(db)
	userRepository := repositories.NewUserRepository(db)

	claims, err := DecodeToken(refreshToken)
	if err != nil {
		return "", nil
	}

	data := claims["data"].(map[string]interface{})
	sessionId := data["session_id"].(string)

	session := &entities.Session{
		Id: sessionId,
	}
	err = sessionRepository.FindSession(session)
	if err != nil {
		return "", nil
	}

	user := entities.User{
		Id: session.UserId,
	}
	_, err = userRepository.FindUser(&user)
	if err != nil {
		return "", nil
	}

	newClaims := &CustomClaim{
		User: &user,
	}

	expireAccessToken := time.Now().Add(time.Hour * 12).UnixMilli()
	accessToken, err := GenerateToken(newClaims, expireAccessToken)
	if err != nil {
		return "", nil
	}

	return accessToken, &user
}
