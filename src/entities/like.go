package entities

type TweetLike struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	TweetId   string `json:"tweet_id" gorm:"size:200;not null" binding:"required"`
	UserId    string `json:"user_id" gorm:"size:200;not null"`
	IsLike    bool   `json:"is_like" gorm:"default:true"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime:milli"`
}
