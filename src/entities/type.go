package entities

type TweetType struct {
	Id   int32  `json:"id" gorm:"primaryKey"`
	Type string `json:"type" gorm:"size:100;not null"`
}
