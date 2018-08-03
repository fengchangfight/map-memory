package entity

import "time"

type User struct {
	ID       int64 `gorm:"primary_key"`
	Username string
	Password string
	Phone    *string
	Nickname string
	Email    string
	RegDate  time.Time
}

func (User) TableName() string {
	return "mp_user"
}
