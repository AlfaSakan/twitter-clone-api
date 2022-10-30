package schemas

type TweetRequest struct {
	Content string `json:"content" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
	TypeId  int32  `json:"type_id" binding:"required"`
}

type TweetRequestById struct {
	Id string `json:"id" binding:"required"`
}

type TweetRequestByUserId struct {
	UserId  string `json:"user_id" binding:"required"`
	Content string `json:"content"`
	Id      string `json:"id"`
}
