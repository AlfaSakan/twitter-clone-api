package repositories

import (
	"github.com/AlfaSakan/twitter-clone-api/src/entities"
	"gorm.io/gorm"
)

type ITweetLikeRepository interface {
	FindLike(like *entities.TweetLike) error
	CreateLike(like *entities.TweetLike) error
	UpdateLike(like *entities.TweetLike) error
	UpdateIsLike(request *entities.TweetLike) error
}

type TweetLikeRepository struct {
	db *gorm.DB
}

func NewTweetLikeRepository(db *gorm.DB) *TweetLikeRepository {
	return &TweetLikeRepository{db}
}

func (r *TweetLikeRepository) CreateLike(like *entities.TweetLike) error {
	return r.db.Create(like).Error
}

func (r *TweetLikeRepository) FindLike(like *entities.TweetLike) error {
	return r.db.Debug().Where(like).First(like).Error
}

func (r *TweetLikeRepository) UpdateLike(like *entities.TweetLike) error {
	query := &entities.TweetLike{
		Id: like.Id,
	}
	return r.db.Where(query).Updates(like).Error
}

func (r *TweetLikeRepository) UpdateIsLike(request *entities.TweetLike) error {
	query := &entities.TweetLike{
		Id: request.Id,
	}

	like := map[string]interface{}{
		"is_like": request.IsLike,
	}

	return r.db.Model(query).Where(query).Updates(like).Error
}
