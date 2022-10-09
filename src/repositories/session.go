package repositories

import (
	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"gorm.io/gorm"
)

type ISessionRepository interface {
	CreateSession(session *entities.Session) error
	FindSession(session *entities.Session) error
	DeleteSession(session *entities.Session) error
}

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *SessionRepository {
	return &SessionRepository{db}
}

func (r *SessionRepository) FindSession(session *entities.Session) error {
	return r.db.Where(&entities.Session{Id: session.Id}).Find(session).Error
}

func (r *SessionRepository) CreateSession(session *entities.Session) error {
	return r.db.Create(session).Error
}

func (r *SessionRepository) DeleteSession(session *entities.Session) error {
	return r.db.Where(&entities.Session{Id: session.Id}).Delete(session).Error
}
