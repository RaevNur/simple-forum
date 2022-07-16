package post

import (
	"fmt"
	"forum/internal/helper/constraints"
	model "forum/internal/models"
)

func (s *PostService) Create(post *model.Post) error {
	err := s.repo.Create(post)
	if err != nil {
		return fmt.Errorf("PostService.Create: %w", err)
	}

	return nil
}

func (s *PostService) GetById(id int64) (*model.Post, error) {
	post, err := s.repo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("PostService.GetByID: %w", err)
	}

	return post, nil
}

func (s *PostService) GetLikesCount(post *model.Post) error {
	err := s.repo.GetLikesCount(post)
	if err != nil {
		return fmt.Errorf("PostService.GetLikesCount: %w", err)
	}

	return nil
}

func (s *PostService) IsLikedByUser(userId int64, post *model.Post) error {
	val, err := s.repo.IsLikedByUser(userId, post.Id)
	if err != nil {
		return fmt.Errorf("PostService.IsLikedByUser: %w", err)
	}

	switch val {
	case constraints.LikeValue:
		post.IsLiked = true
	case constraints.DislikeValue:
		post.IsDisliked = true
	case 0:
		return nil
	default:
		return fmt.Errorf("PostService.IsLikedByUser: %w", &constraints.ValidateError{
			Field:       "like value",
			Description: "like value is not same as constraints",
		})
	}

	return nil
}

// func (s *PostService) GetByLiked(userId int64, page int) ([]*model.Post, error) {
// 	posts, err := s.repo.GetByLiked(userId, page)
// 	if err != nil {
// 		return nil, fmt.Errorf("PostService.GetByLiked: %w", err)
// 	}

// 	return posts, nil
// }
