package tag

import (
	"fmt"
	model "forum/internal/models"
)

func (s *TagService) Create(threadId int64, tag string) error {
	t, err := s.repo.GetTagByName(tag)
	if err != nil {
		return fmt.Errorf("TagService.Create: %w", err)
	}

	if t == nil {
		t = &model.Tag{
			Name: tag,
		}

		err := s.repo.Create(t)
		if err != nil {
			return fmt.Errorf("TagService.Create: %w", err)
		}
	}

	err = s.repo.CreateRelation(threadId, t)
	if err != nil {
		return fmt.Errorf("TagService.Create: %w", err)
	}

	return nil
}

func (s *TagService) CreateTags(threadId int64, tags []string) error {
	for _, tag := range tags {
		t, err := s.repo.GetTagByName(tag)
		if err != nil {
			return fmt.Errorf("TagService.Create: %w", err)
		}

		if t == nil {
			t = &model.Tag{
				Name: tag,
			}

			err := s.repo.Create(t)
			if err != nil {
				return fmt.Errorf("TagService.Create: %w", err)
			}
		}

		err = s.repo.CreateRelation(threadId, t)
		if err != nil {
			return fmt.Errorf("TagService.Create: %w", err)
		}
	}

	return nil
}

func (s *TagService) GetTags() ([]*model.Tag, error) {
	tags, err := s.repo.GetTags()
	if err != nil {
		return nil, fmt.Errorf("TagService.GetTags: %w", err)
	}

	return tags, nil
}

func (s *TagService) GetTagByName(name string) (*model.Tag, error) {
	tag, err := s.repo.GetTagByName(name)
	if err != nil {
		return nil, fmt.Errorf("TagService.GetTagByName: %w", err)
	}

	return tag, nil
}

// for future pagination
// func (s *TagService) GetTags(page int) ([]*model.Tag, error) {
// 	tags, err := s.repo.GetTags(page)
// 	if err != nil {
// 		return nil, fmt.Errorf("TagService.GetTags: %w", err)
// 	}

// 	return tags, nil
// }

func (s *TagService) GetTagsByThread(threadId int64) ([]*model.Tag, error) {
	tags, err := s.repo.GetTagsByThread(threadId)
	if err != nil {
		return nil, fmt.Errorf("TagService.GetTagsByThread: %w", err)
	}

	return tags, nil
}
