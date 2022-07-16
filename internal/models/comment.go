package models

type Comment struct {
	Id       int64
	Author   *User
	Post     *Post
	ThreadId int64
}

type ICommentRepo interface {
	Create(comment *Comment) error
	GetCommentsByThread(threadId int64) ([]*Comment, error)
}

type ICommentService interface {
	Comment(comment *Comment) error
	GetThreadComments(threadId int64) ([]*Comment, error)
}
