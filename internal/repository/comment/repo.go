package comment

import (
	"fmt"
	"forum/internal/helper"
	"forum/internal/helper/constraints"

	model "forum/internal/models"
)

func (c *CommentRepo) Create(comment *model.Comment) error {
	query := `INSERT INTO comments (
		post_id,
		thread_id
	) 
	VALUES (?, ?);`

	stmt, err := c.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("CommentRepo.Create: %w", err)
	}

	defer stmt.Close()

	res, err := stmt.Exec((*comment).Post.Id, (*comment).ThreadId)
	if err != nil {
		return fmt.Errorf("CommentRepo.Create: %w", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("CommentRepo.Create: %w", err)
	}

	(*comment).Id = lastId
	return nil
}

// ordered by created time (asc)
func (c *CommentRepo) GetCommentsByThread(threadId int64) ([]*model.Comment, error) {
	query := `SELECT comments.id, comments.thread_id, posts.id, posts.content, posts.user_id, posts.created_at FROM comments 
	INNER JOIN posts ON posts.id = comments.post_id
	WHERE comments.thread_id = ?
	ORDER BY posts.created_at ASC`

	rows, err := c.db.Query(query, threadId)
	if err != nil {
		return nil, fmt.Errorf("CommentRepo.GetCommentsByThread: %w", err)
	}

	comments := make([]*model.Comment, 0)
	for rows.Next() {
		t := model.Comment{}
		p := model.Post{}
		var decodedTime string

		err = rows.Scan(
			&t.Id,
			&t.ThreadId,
			&p.Id,
			&p.Content,
			&p.UserId,
			&decodedTime,
		)
		if err != nil {
			return nil, fmt.Errorf("CommentRepo.GetCommentsByThread: %w", err)
		}

		p.CreatedAt, err = helper.DecodeTime(decodedTime, constraints.TimeFormatRFC3339)
		if err != nil {
			return nil, fmt.Errorf("CommentRepo.GetCommentsByThread: %w", err)
		}

		t.Post = &p
		comments = append(comments, &t)
	}

	return comments, nil
}
