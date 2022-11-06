package schemas

type PostRetweetSchema struct {
	Content string `json:"content"`
	TweetId string `json:"tweet_id" binding:"required"`
}
