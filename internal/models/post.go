package models

import (
	"forum/internal/helper/constraints"
	"time"
)

type Post struct {
	Id        int64
	Content   string
	UserId    int64
	CreatedAt time.Time
	Likes     int
	Dislikes  int
	// Shows if it's liked/disliked bu current user
	IsLiked    bool
	IsDisliked bool
}

func (p *Post) TimeView() string {
	return p.CreatedAt.Format(constraints.TimeFormatView)
}

type IPostRepo interface {
	Create(post *Post) error
	GetById(id int64) (*Post, error)
	GetLikesCount(post *Post) error
	IsLikedByUser(userId, postId int64) (int, error)
	// GetByLiked(userId int64, page int) ([]*Post, error)
}

type IPostService interface {
	Create(post *Post) error
	GetById(id int64) (*Post, error)
	GetLikesCount(post *Post) error
	IsLikedByUser(userId int64, post *Post) error
	// GetByLiked(userId int64, page int) ([]*Post, error)
}
