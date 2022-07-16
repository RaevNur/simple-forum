package repository

import (
	"database/sql"
	"forum/internal/models"
	"forum/internal/repository/comment"
	"forum/internal/repository/like"
	"forum/internal/repository/post"
	"forum/internal/repository/session"
	"forum/internal/repository/tag"
	"forum/internal/repository/thread"
	"forum/internal/repository/user"
)

type Repository struct {
	User    models.IUserRepo
	Session models.ISessionRepo
	Thread  models.IThreadRepo
	Post    models.IPostRepo
	Comment models.ICommentRepo
	Like    models.ILikeRepo
	Tag     models.ITagRepo
}

func NewRepo(db *sql.DB) *Repository {
	return &Repository{
		User:    user.NewUserRepo(db),
		Session: session.NewSessionRepo(db),
		Thread:  thread.NewThreadRepo(db),
		Post:    post.NewPostRepo(db),
		Comment: comment.NewCommentRepo(db),
		Like:    like.NewLikeRepo(db),
		Tag:     tag.NewTagRepo(db),
	}
}
