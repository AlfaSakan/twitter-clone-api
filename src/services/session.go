package services

import (
	"net/http"
	"time"

	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"github.com/AlfaSakan/twitter-clone-api/src/helpers"
	"github.com/AlfaSakan/twitter-clone-api/src/repositories"
	"github.com/AlfaSakan/twitter-clone-api/src/schemas"
	"github.com/AlfaSakan/twitter-clone-api/src/utils"
	"golang.org/x/crypto/bcrypt"
)

type ISessionService interface {
	Login(request *schemas.SessionRequest, user *entities.User, session *entities.Session) (statusCode int, errorMessage error)
	Logout(request *entities.Session) error
	GenerateAccessRefresh(user *entities.User, session *entities.Session) (string, string, error)
}

type SessionService struct {
	sessionRepository repositories.ISessionRepository
	userRepository    repositories.IUserRepository
}

func NewSessionService(
	sessionRepository repositories.ISessionRepository,
	userRepository repositories.IUserRepository,
) *SessionService {
	return &SessionService{sessionRepository, userRepository}
}

func (s *SessionService) Login(request *schemas.SessionRequest, user *entities.User, session *entities.Session) (int, error) {
	user.Username = request.Username

	status, err := s.userRepository.FindUser(user)
	if err != nil {
		return status, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return http.StatusBadRequest, helpers.WrongPassword
	}

	session.Id = helpers.GenerateId()
	session.Valid = true
	session.UserAgent = request.UserAgent
	session.UserId = user.Id

	return http.StatusCreated, s.sessionRepository.CreateSession(session)
}

func (s *SessionService) GenerateAccessRefresh(user *entities.User, session *entities.Session) (string, string, error) {
	accessClaims := &utils.CustomClaim{
		User:      user,
		SessionId: session.Id,
	}
	expireAccessToken := time.Now().Add(time.Hour * 12).UnixMilli()
	accessToken, err := utils.GenerateToken(accessClaims, expireAccessToken)
	if err != nil {
		return "", "", err
	}

	refreshClaim := &utils.CustomClaim{
		User:      user,
		SessionId: session.Id,
	}

	expireRefreshToken := time.Now().Add(time.Hour * 24 * 30 * 12).UnixMilli()
	refreshToken, err := utils.GenerateToken(refreshClaim, expireRefreshToken)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *SessionService) Logout(request *entities.Session) error {
	return s.sessionRepository.DeleteSession(request)
}
