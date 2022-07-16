package session

import "forum/internal/models"

type SessionService struct {
	repo models.ISessionRepo
}

func NewSessionService(repo models.ISessionRepo) *SessionService {
	return &SessionService{repo}
}
