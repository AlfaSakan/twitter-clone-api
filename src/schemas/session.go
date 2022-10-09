package schemas

type SessionRequest struct {
	UserAgent string `json:"user_agent"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
