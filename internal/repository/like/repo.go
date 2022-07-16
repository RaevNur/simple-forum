package like

import (
	"database/sql"
	"fmt"

	model "forum/internal/models"
)

func (l *LikeRepo) Create(like *model.Like) error {
	query := `INSERT INTO likes (
		user_id, 
		post_id, 
		liked
	) 
	VALUES (?, ?, ?);`

	stmt, err := l.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("LikeRepo.Create: %w", err)
	}

	defer stmt.Close()

	res, err := stmt.Exec((*like).UserId, (*like).PostId, (*like).Liked)
	if err != nil {
		return fmt.Errorf("LikeRepo.Create: %w", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("LikeRepo.Create: %w", err)
	}

	(*like).Id = lastId
	return nil
}

// updates only 'liked' value
func (l *LikeRepo) Update(like *model.Like) error {
	query := `UPDATE likes SET liked = ? WHERE id = ?`

	stmt, err := l.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("LikeRepo.Update: %w", err)
	}

	defer stmt.Close()

	res, err := stmt.Exec((*like).Liked, (*like).Id)
	if err != nil {
		return fmt.Errorf("LikeRepo.Update: %w", err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("LikeRepo.Update: %w", err)
	}
	if affect != 1 {
		return fmt.Errorf("LikeRepo.Update affected rows more than 1: %d", affect)
	}

	return nil
}

func (l *LikeRepo) Exist(like *model.Like) (bool, error) {
	query := `SELECT id FROM likes WHERE user_id = ? AND post_id = ?`
	row := l.db.QueryRow(query, like.UserId, like.PostId)

	err := row.Scan(&like.Id)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("LikeRepo.Exist: %w", err)
	}

	return true, nil
}

func (l *LikeRepo) Delete(id int64) error {
	query := `DELETE FROM likes WHERE id = ?`

	stmt, err := l.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("LikeRepo.Delete: %w", err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("LikeRepo.Delete: %w", err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("LikeRepo.Delete: %w", err)
	}
	if affect != 1 {
		return fmt.Errorf("LikeRepo.Delete affected rows more than 1: %d", affect)
	}

	return nil
}
