package thread

import (
	"fmt"
	model "forum/internal/models"
)

func (s *ThreadService) Create(thread *model.Thread) error {
	err := s.repo.Create(thread)
	if err != nil {
		return fmt.Errorf("ThreadService.Create: %w", err)
	}

	return nil
}

func (s *ThreadService) GetById(id int64) (*model.Thread, error) {
	thread, err := s.repo.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("ThreadService.GetById: %w", err)
	}

	return thread, nil
}

func (s *ThreadService) GetUserCreatedThreads(userId int64) ([]*model.Thread, error) {
	threads, err := s.repo.GetByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf("ThreadService.GetUserCreatedThreads: %w", err)
	}

	return threads, nil
}

func (s *ThreadService) GetUserLikedThreads(userId int64) ([]*model.Thread, error) {
	threads, err := s.repo.GetByLiked(userId)
	if err != nil {
		return nil, fmt.Errorf("ThreadService.GetUserLikedThreads: %w", err)
	}

	return threads, nil
}

func (s *ThreadService) GetByTag(tagId int64) ([]*model.Thread, error) {
	threads, err := s.repo.GetByTag(tagId)
	if err != nil {
		return nil, fmt.Errorf("ThreadService.GetByTag: %w", err)
	}

	return threads, nil
}

func (s *ThreadService) GetRecentQuestions() ([]*model.Thread, error) {
	threads, err := s.repo.GetRecentQuestions()
	if err != nil {
		return nil, fmt.Errorf("ThreadService.GetRecentQuestions: %w", err)
	}

	return threads, nil
}

// func (s *ThreadService) GetRecentQuestions(page int) ([]*model.Thread, error) {
// 	threads, err := s.repo.GetRecentQuestions(page)
// 	if err != nil {
// 		return nil, fmt.Errorf("ThreadService.GetRecentQuestions: %w", err)
// 	}

// 	return threads, nil
// }

// func (s *ThreadService) GetUserCreatedThreads(userId int64, page int) ([]*model.Thread, error) {
// 	threads, err := s.repo.GetByUserId(userId, page)
// 	if err != nil {
// 		return nil, fmt.Errorf("ThreadService.GetUserCreatedThreads: %w", err)
// 	}

// 	return threads, nil
// }

// func (s *ThreadService) GetUserLikedThreads(userId int64, page int) ([]*model.Thread, error) {
// 	threads, err := s.repo.GetByLiked(userId, page)
// 	if err != nil {
// 		return nil, fmt.Errorf("ThreadService.GetUserLikedThreads: %w", err)
// 	}

// 	return threads, nil
// }

// func (s *ThreadService) GetByTag(tagId int64, page int) ([]*model.Thread, error) {
// 	threads, err := s.repo.GetByTag(tagId, page)
// 	if err != nil {
// 		return nil, fmt.Errorf("ThreadService.GetByTag: %w", err)
// 	}

// 	return threads, nil
// }
