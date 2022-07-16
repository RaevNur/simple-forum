package models

type Like struct {
	Id     int64
	UserId int64
	PostId int64
	Liked  int
}

type ILikeRepo interface {
	Create(like *Like) error
	Update(like *Like) error
	Exist(like *Like) (bool, error)
	Delete(id int64) error
}

type ILikeService interface {
	Like(like *Like) error
	Dislike(like *Like) error
	Unlike(like *Like) error
}
