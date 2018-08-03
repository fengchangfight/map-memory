package model

type UserVO struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
