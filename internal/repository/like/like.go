package like

import "database/sql"

type LikeRepo struct {
	db *sql.DB
}

func NewLikeRepo(db *sql.DB) *LikeRepo {
	return &LikeRepo{db}
}
