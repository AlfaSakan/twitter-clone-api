package schemas

type UserRequest struct {
	Id       string `json:"id" binding:"required"`
	Username string `jsons:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
