package tag

import "forum/internal/models"

type TagService struct {
	repo models.ITagRepo
}

func NewTagService(repo models.ITagRepo) *TagService {
	return &TagService{repo}
}
