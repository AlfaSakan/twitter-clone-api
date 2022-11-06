package entities

type Retweet struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	TweetId   string `json:"tweet_id" gorm:"size:200;not null"`
	RetweetId string `json:"retweet_id" gorm:"size:200;not null"`
}
