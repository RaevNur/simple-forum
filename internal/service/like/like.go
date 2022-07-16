package like

import "forum/internal/models"

type LikeService struct {
	repo models.ILikeRepo
}

func NewLikeService(repo models.ILikeRepo) *LikeService {
	return &LikeService{repo}
}
