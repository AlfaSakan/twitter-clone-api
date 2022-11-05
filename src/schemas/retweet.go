package schemas

type PostRetweetSchema struct {
	Content string `json:"content" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
	TypeId  int32  `json:"type_id" binding:"required"`
	TweetId string `json:"tweet_id" binding:"required"`
}
