package services

import (
	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"github.com/AlfaSakan/twitter-clone-api/src/repositories"
)

type ISessionService interface {
	Login(request *entities.Session) error
	Logout(request *entities.Session) error
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

func (s *SessionService) Login(request *entities.Session) error {
	return s.sessionRepository.CreateSession(request)

}

func (s *SessionService) Logout(request *entities.Session) error {
	return s.sessionRepository.DeleteSession(request)
}
