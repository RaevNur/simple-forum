package service

import (
	"forum/internal/models"
	"forum/internal/repository"
	"forum/internal/service/comment"
	"forum/internal/service/like"
	"forum/internal/service/post"
	"forum/internal/service/session"
	"forum/internal/service/tag"
	"forum/internal/service/thread"
	"forum/internal/service/user"
)

type Service struct {
	User    models.IUserService
	Thread  models.IThreadService
	Post    models.IPostService
	Comment models.ICommentService
	Session models.ISessionService
	Tag     models.ITagService
	Like    models.ILikeService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:    user.NewUserService(repo.User),
		Thread:  thread.NewThreadService(repo.Thread),
		Session: session.NewSessionService(repo.Session),
		Post:    post.NewPostService(repo.Post),
		Comment: comment.NewCommentService(repo.Comment),
		Tag:     tag.NewTagService(repo.Tag),
		Like:    like.NewLikeService(repo.Like),
	}
}
