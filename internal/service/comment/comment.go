package comment

import "forum/internal/models"

type CommentService struct {
	repo models.ICommentRepo
}

func NewCommentService(repo models.ICommentRepo) *CommentService {
	return &CommentService{repo}
}
