package entities

type Tweet struct {
	Id        string `json:"id" gorm:"primaryKey;size:200"`
	Content   string `json:"content" gorm:"size:250;not null" binding:"required"`
	UserId    string `json:"user_id" gorm:"size:200;not null" binding:"required"`
	Likes     int32  `json:"likes" gorm:"default:0"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime:milli"`
	IsLike    bool   `json:"is_like"`
	User      User   `gorm:"foreignKey:UserId;references:id"`
}