package entities

type TweetReply struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	TweetId string `json:"tweet_id" gorm:"size:200;not null"`
	ReplyId string `json:"user_id" gorm:"size:200;not null"`
}

const (
	_ int32 = iota
	TypeTweet
	TypeRetweet
	TypeReply
)
