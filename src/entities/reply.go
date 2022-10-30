package entities

type TweetReply struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	TweetId string `json:"tweet_id" gorm:"size:200;not null"`
	ReplyId string `json:"user_id" gorm:"size:200;not null"`
}

const (
	TypeTweet   int32 = 1
	TypeRetweet int32 = 2
	TypeReply   int32 = 3
)
