package models

import (
	"forum/internal/helper/constraints"
	"time"
)

// User -
type User struct {
	Id        int64
	Nickname  string
	Firstname string
	Lastname  string
	Email     string
	Password  string
	CreatedAt time.Time
	// Avatar    string
}

func (u *User) TimeView() string {
	return u.CreatedAt.Format(constraints.TimeFormatView)
}

type IUserRepo interface {
	Create(user *User) error
	GetByID(id int64) (*User, error)
	GetPassword(nickname, email string) (*User, error)
	GetByNickname(nickname string) (*User, error)
	UserExist(nickname, email string) (bool, error)
}

type IUserService interface {
	Register(user *User) error
	Login(user *User) error
	GetByID(id int64) (*User, error)
	GetByNickname(nickname string) (*User, error)
}
