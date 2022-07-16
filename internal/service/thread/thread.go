package thread

import "forum/internal/models"

type ThreadService struct {
	repo models.IThreadRepo
}

func NewThreadService(repo models.IThreadRepo) *ThreadService {
	return &ThreadService{repo}
}
