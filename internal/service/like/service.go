package like

import (
	"fmt"
	"forum/internal/helper/constraints"
	model "forum/internal/models"
)

func (s *LikeService) Like(like *model.Like) error {
	exist, err := s.repo.Exist(like)
	if err != nil {
		return fmt.Errorf("LikeService.Like: %w", err)
	}

	if exist {
		err = s.repo.Update(like)
		if err != nil {
			return fmt.Errorf("LikeService.Like: %w", err)
		}
	} else {
		err = s.repo.Create(like)
		if err != nil {
			return fmt.Errorf("LikeService.Like: %w", err)
		}
	}

	return nil
}

func (s *LikeService) Dislike(like *model.Like) error {
	exist, err := s.repo.Exist(like)
	if err != nil {
		return fmt.Errorf("LikeService.Dislike: %w", err)
	}

	if exist {
		err = s.repo.Update(like)
		if err != nil {
			return fmt.Errorf("LikeService.Dislike: %w", err)
		}
	} else {
		err = s.repo.Create(like)
		if err != nil {
			return fmt.Errorf("LikeService.Dislike: %w", err)
		}
	}

	return nil
}

func (s *LikeService) Unlike(like *model.Like) error {
	exist, err := s.repo.Exist(like)
	if err != nil {
		return fmt.Errorf("LikeService.Unlike: %w", err)
	}

	if !exist {
		return fmt.Errorf("LikeService.Unlike: %w", &constraints.ExistsError{
			Title:       "like",
			Description: "cant find like/dislike by user on post",
		})
	}

	err = s.repo.Delete(like.Id)
	if err != nil {
		return fmt.Errorf("LikeService.Unlike: %w", err)
	}

	return nil
}
