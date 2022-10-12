package entities

type User struct {
	Id        string `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"size:50;unique;not null" binding:"required"`
	Name      string `json:"name" gorm:"size:100;not null"`
	Email     string `json:"email" gorm:"size:100;not null"`
	Password  string `json:"password" gorm:"size:200;not null" binding:"required"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoCreateTime:milli"`
}
