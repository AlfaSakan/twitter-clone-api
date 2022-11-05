package entities

type Tweet struct {
	Id            string `json:"id" gorm:"primaryKey;size:200"`
	Content       string `json:"content" gorm:"size:250;not null" binding:"required"`
	UserId        string `json:"user_id" gorm:"size:200;not null" binding:"required"`
	Likes         int32  `json:"likes" gorm:"default:0"`
	ReplyCounts   int32  `json:"reply_counts" gorm:"default:0"`
	RetweetCounts int32  `json:"retweet_counts" gorm:"default:0"`
	TypeId        int32  `json:"type_id" gorm:"not null"`
	CreatedAt     int64  `json:"created_at" gorm:"autoCreateTime:milli"`
	User          User   `gorm:"foreignKey:UserId;references:id"`

	IsLike bool   `json:"is_like" gorm:"-"`
	Type   string `json:"type" gorm:"-"`
}
