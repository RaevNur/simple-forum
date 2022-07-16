package models

import "time"

type Session struct {
	Id        int64
	Uuid      string
	UserId    int64
	CreatedAt time.Time
}

type ISessionRepo interface {
	Create(session *Session) error
	Delete(id int64) error
	GetByUserId(userId int64) (*Session, error)
	GetByUuid(uuid string) (*Session, error)
}

type ISessionService interface {
	GenerateSession(userId int64) (*Session, error)
	DeleteSession(id int64) error
	GetByUserId(userId int64) (*Session, error)
	GetByUuid(uuid string) (*Session, error)
}
