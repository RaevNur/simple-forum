package user

import "forum/internal/models"

type UserService struct {
	repo models.IUserRepo
}

func NewUserService(repo models.IUserRepo) *UserService {
	return &UserService{repo}
}
