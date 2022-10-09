package entities

type Session struct {
	Id        string `json:"id" gorm:"primaryKey"`
	Valid     bool   `json:"valid" gorm:"default:true"`
	UserAgent string `json:"user_agent" gorm:"size:200;not null"`
	UserId    string `json:"user_id" gorm:"not null"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	SessionId    string `json:"session_id"`
}
