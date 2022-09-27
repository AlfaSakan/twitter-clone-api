package schemas

type PostUser struct {
	Username string `jsons:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
}
