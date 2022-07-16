package thread

import "database/sql"

type ThreadRepo struct {
	db *sql.DB
}

func NewThreadRepo(db *sql.DB) *ThreadRepo {
	return &ThreadRepo{db}
}
