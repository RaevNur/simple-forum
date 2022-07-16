package comment

import (
	"fmt"
	model "forum/internal/models"
)

func (s *CommentService) Comment(comment *model.Comment) error {
	err := s.repo.Create(comment)
	if err != nil {
		return fmt.Errorf("CommentService.Comment: %w", err)
	}

	return nil
}

func (s *CommentService) GetThreadComments(threadId int64) ([]*model.Comment, error) {
	comments, err := s.repo.GetCommentsByThread(threadId)
	if err != nil {
		return nil, fmt.Errorf("CommentService.GetThreadComments: %w", err)
	}

	return comments, nil
}
