package model

type CreateUserRequest struct {
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
