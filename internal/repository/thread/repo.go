package thread

import (
	"database/sql"
	"fmt"
	"forum/internal/helper"
	"forum/internal/helper/constraints"

	model "forum/internal/models"
)

func (t *ThreadRepo) Create(thread *model.Thread) error {
	query := `INSERT INTO threads (
		post_id, 
		title
	) 
	VALUES (?, ?);`

	stmt, err := t.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("ThreadRepo.Create: %w", err)
	}

	defer stmt.Close()

	res, err := stmt.Exec((*thread).Post.Id, (*thread).Title)
	if err != nil {
		return fmt.Errorf("ThreadRepo.Create: %w", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("ThreadRepo.Create: %w", err)
	}

	(*thread).Id = lastId
	return nil
}

// fills the post
func (t *ThreadRepo) GetById(id int64) (*model.Thread, error) {
	query := `SELECT threads.title, posts.id, posts.content, posts.user_id, posts.created_at FROM threads
	INNER JOIN posts ON threads.post_id = posts.id 
	WHERE threads.id = ?`
	row := t.db.QueryRow(query, id)

	thread := model.Thread{
		Id: id,
	}
	post := model.Post{}
	var encodedTime string

	err := row.Scan(&thread.Title, &post.Id, &post.Content, &post.UserId, &encodedTime)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("ThreadRepo.GetById: %w", err)
	}

	post.CreatedAt, err = helper.DecodeTime(encodedTime, constraints.TimeFormatRFC3339)
	if err != nil {
		return nil, fmt.Errorf("ThreadRepo.GetById: %w", err)
	}

	thread.Post = &post
	return &thread, nil
}

// also fills post
func (t *ThreadRepo) GetByLiked(userId int64) ([]*model.Thread, error) {
	query := `SELECT threads.id, threads.title, posts.id, posts.content, posts.created_at FROM threads 
	INNER JOIN posts ON threads.post_id = posts.id 
	INNER JOIN likes ON posts.id = likes.post_id
	WHERE likes.user_id = ? AND likes.liked = ?
	ORDER BY posts.created_at DESC`

	rows, err := t.db.Query(query, userId, constraints.LikeValue)
	if err != nil {
		return nil, fmt.Errorf("ThreadRepo.GetByLiked: %w", err)
	}

	threads := make([]*model.Thread, 0)
	for rows.Next() {
		t := model.Thread{}
		p := model.Post{}
		var encodedTime string

		err = rows.Scan(
			&t.Id,
			&t.Title,
			&p.Id,
			&p.Content,
			&encodedTime,
		)
		if err != nil {
			return nil, fmt.Errorf("ThreadRepo.GetByLiked: %w", err)
		}

		p.CreatedAt, err = helper.DecodeTime(encodedTime, constraints.TimeFormatRFC3339)
		if err != nil {
			return nil, fmt.Errorf("ThreadRepo.GetByLiked: %w", err)
		}

		p.UserId = userId
		t.Post = &p
		threads = append(threads, &t)
	}

	return threads, nil
}

// takes a page number (1, 2, 3...)
// also fills post
// func (t *ThreadRepo) GetByLiked(userId int64, page int) ([]*model.Thread, error) {
// 	query := `SELECT threads.id, threads.title, posts.id, posts.content, posts.created_at FROM threads
// 	INNER JOIN posts ON threads.post_id = posts.id
// 	INNER JOIN likes ON post.id = likes.post_id
// 	WHERE likes.user_id = ? AND likes.liked = ?
// 	ORDER BY posts.created_at DESC
// 	LIMIT ? OFFSET ?`

// 	offset := (page - 1) * constraints.LimitThreadsPerPage
// 	rows, err := t.db.Query(query, userId, constraints.LikeValue, constraints.LimitThreadsPerPage, offset)
// 	if err != nil {
// 		return nil, fmt.Errorf("ThreadRepo.GetByLiked: %w", err)
// 	}

// 	threads := make([]*model.Thread, 0, constraints.LimitThreadsPerPage)
// 	for rows.Next() {
// 		t := model.Thread{}
// 		p := model.Post{}
// 		var encodedTime string

// 		err = rows.Scan(
// 			&t.Id,
// 			&t.Title,
// 			&p.Id,
// 			&p.Content,
// 			&encodedTime,
// 		)
// 		if err != nil {
// 			return nil, fmt.Errorf("ThreadRepo.GetByLiked: %w", err)
// 		}

// 		p.CreatedAt, err = helper.DecodeTime(encodedTime, constraints.TimeFormatRFC3339)
// 		if err != nil {
// 			return nil, fmt.Errorf("ThreadRepo.GetByLiked: %w", err)
// 		}

// 		p.UserId = userId
// 		t.Post = &p
// 		threads = append(threads, &t)
// 	}

// 	return threads, nil
// }

// also fills post
func (t *ThreadRepo) GetByUserId(userId int64) ([]*model.Thread, error) {
	query := `SELECT threads.id, threads.title, posts.id, posts.content, posts.created_at FROM threads 
	INNER JOIN posts ON threads.post_id = posts.id 
	WHERE posts.user_id = ? 
	ORDER BY posts.created_at DESC`

	rows, err := t.db.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("ThreadRepo.GetByUserId: %w", err)
	}

	threads := make([]*model.Thread, 0)
	for rows.Next() {
		t := model.Thread{}
		p := model.Post{}
		var encodedTime string

		err = rows.Scan(
			&t.Id,
			&t.Title,
			&p.Id,
			&p.Content,
			&encodedTime,
		)
		if err != nil {
			return nil, fmt.Errorf("ThreadRepo.GetByUserId: %w", err)
		}

		p.CreatedAt, err = helper.DecodeTime(encodedTime, constraints.TimeFormatRFC3339)
		if err != nil {
			return nil, fmt.Errorf("ThreadRepo.GetByUserId: %w", err)
		}

		p.UserId = userId
		t.Post = &p
		threads = append(threads, &t)
	}

	return threads, nil
}

// takes a page number (1, 2, 3...)
// also fills post
// func (t *ThreadRepo) GetByUserId(userId int64, page int) ([]*model.Thread, error) {
// 	query := `SELECT threads.id, threads.title, posts.id, posts.content, posts.created_at FROM threads
// 	INNER JOIN posts ON threads.post_id = posts.id
// 	WHERE posts.user_id = ?
// 	ORDER BY posts.created_at DESC
// 	LIMIT ? OFFSET ?`

// 	offset := (page - 1) * constraints.LimitThreadsPerPage
// 	rows, err := t.db.Query(query, userId, constraints.LimitThreadsPerPage, offset)
// 	if err != nil {
// 		return nil, fmt.Errorf("ThreadRepo.GetByUserId: %w", err)
// 	}

// 	threads := make([]*model.Thread, 0, constraints.LimitThreadsPerPage)
// 	for rows.Next() {
// 		t := model.Thread{}
// 		p := model.Post{}
// 		var encodedTime string

// 		err = rows.Scan(
// 			&t.Id,
// 			&t.Title,
// 			&p.Id,
// 			&p.Content,
// 			&encodedTime,
// 		)
// 		if err != nil {
// 			return nil, fmt.Errorf("ThreadRepo.GetByUserId: %w", err)
// 		}

// 		p.CreatedAt, err = helper.DecodeTime(encodedTime, constraints.TimeFormatRFC3339)
// 		if err != nil {
// 			return nil, fmt.Errorf("ThreadRepo.GetByUserId: %w", err)
// 		}

// 		p.UserId = userId
// 		t.Post = &p
// 		threads = append(threads, &t)
// 	}

// 	return threads, nil
// }

// also fills post
func (t *ThreadRepo) GetByTag(tagId int64) ([]*model.Thread, error) {
	query := `SELECT threads.id, threads.title, posts.id, posts.content, posts.user_id, posts.created_at FROM threads 
	INNER JOIN posts ON threads.post_id = posts.id 
	INNER JOIN tags_threads AS tt ON threads.id = tt.thread_id
	WHERE tt.tag_id = ? 
	ORDER BY posts.created_at DESC`

	rows, err := t.db.Query(query, tagId)
	if err != nil {
		return nil, fmt.Errorf("ThreadRepo.GetByTag: %w", err)
	}

	threads := make([]*model.Thread, 0)
	for rows.Next() {
		t := model.Thread{}
		p := model.Post{}
		var encodedTime string

		err = rows.Scan(
			&t.Id,
			&t.Title,
			&p.Id,
			&p.Content,
			&p.UserId,
			&encodedTime,
		)
		if err != nil {
			return nil, fmt.Errorf("ThreadRepo.GetByTag: %w", err)
		}

		p.CreatedAt, err = helper.DecodeTime(encodedTime, constraints.TimeFormatRFC3339)
		if err != nil {
			return nil, fmt.Errorf("ThreadRepo.GetByTag: %w", err)
		}

		t.Post = &p
		threads = append(threads, &t)
	}

	return threads, nil
}

// takes a page number (1, 2, 3...)
// also fills post
// func (t *ThreadRepo) GetByTag(tagId int64, page int) ([]*model.Thread, error) {
// 	query := `SELECT threads.id, threads.title, posts.id, posts.content, posts.user_id, posts.created_at FROM threads
// 	INNER JOIN posts ON threads.post_id = posts.id
// 	INNER JOIN tags_threads AS tt ON threads.id = tt.thread_id
// 	WHERE tt.tag_id = ?
// 	ORDER BY posts.created_at DESC
// 	LIMIT ? OFFSET ?`

// 	offset := (page - 1) * constraints.LimitThreadsPerPage
// 	rows, err := t.db.Query(query, tagId, constraints.LimitThreadsPerPage, offset)
// 	if err != nil {
// 		return nil, fmt.Errorf("ThreadRepo.GetByTag: %w", err)
// 	}

// 	threads := make([]*model.Thread, 0, constraints.LimitThreadsPerPage)
// 	for rows.Next() {
// 		t := model.Thread{}
// 		p := model.Post{}
// 		var encodedTime string

// 		err = rows.Scan(
// 			&t.Id,
// 			&t.Title,
// 			&p.Id,
// 			&p.Content,
// 			&p.UserId,
// 			&encodedTime,
// 		)
// 		if err != nil {
// 			return nil, fmt.Errorf("ThreadRepo.GetByTag: %w", err)
// 		}

// 		p.CreatedAt, err = helper.DecodeTime(encodedTime, constraints.TimeFormatRFC3339)
// 		if err != nil {
// 			return nil, fmt.Errorf("ThreadRepo.GetByTag: %w", err)
// 		}

// 		t.Post = &p
// 		threads = append(threads, &t)
// 	}

// 	return threads, nil
// }

// also fills posts
func (t *ThreadRepo) GetRecentQuestions() ([]*model.Thread, error) {
	query := `SELECT threads.id, threads.title, posts.id, posts.content, posts.user_id, posts.created_at FROM threads 
	INNER JOIN posts ON threads.post_id = posts.id
	ORDER BY posts.created_at DESC`

	rows, err := t.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("ThreadRepo.GetRecentQuestions: %w", err)
	}

	threads := make([]*model.Thread, 0)
	for rows.Next() {
		t := model.Thread{}
		p := model.Post{}
		var encodedTime string

		err = rows.Scan(
			&t.Id,
			&t.Title,
			&p.Id,
			&p.Content,
			&p.UserId,
			&encodedTime,
		)
		if err != nil {
			return nil, fmt.Errorf("ThreadRepo.GetRecentQuestions: %w", err)
		}

		p.CreatedAt, err = helper.DecodeTime(encodedTime, constraints.TimeFormatRFC3339)
		if err != nil {
			return nil, fmt.Errorf("ThreadRepo.GetRecentQuestions: %w", err)
		}

		t.Post = &p
		threads = append(threads, &t)
	}

	return threads, nil
}

// takes a page number (1, 2, 3...)
// also fills post
// func (t *ThreadRepo) GetRecentQuestions(page int) ([]*model.Thread, error) {
// 	query := `SELECT threads.id, threads.title, posts.id, posts.content, posts.user_id, posts.created_at FROM threads
// 	INNER JOIN posts ON threads.post_id = posts.id
// 	ORDER BY posts.created_at DESC
// 	LIMIT ? OFFSET ?`

// 	offset := (page - 1) * constraints.LimitThreadsPerPage
// 	rows, err := t.db.Query(query, constraints.LimitThreadsPerPage, offset)
// 	if err != nil {
// 		return nil, fmt.Errorf("ThreadRepo.GetRecentQuestions: %w", err)
// 	}

// 	threads := make([]*model.Thread, 0, constraints.LimitThreadsPerPage)
// 	for rows.Next() {
// 		t := model.Thread{}
// 		p := model.Post{}
// 		var encodedTime string

// 		err = rows.Scan(
// 			&t.Id,
// 			&t.Title,
// 			&p.Id,
// 			&p.Content,
// 			&p.UserId,
// 			&encodedTime,
// 		)
// 		if err != nil {
// 			return nil, fmt.Errorf("ThreadRepo.GetRecentQuestions: %w", err)
// 		}

// 		p.CreatedAt, err = helper.DecodeTime(encodedTime, constraints.TimeFormatRFC3339)
// 		if err != nil {
// 			return nil, fmt.Errorf("ThreadRepo.GetRecentQuestions: %w", err)
// 		}

// 		t.Post = &p
// 		threads = append(threads, &t)
// 	}

// 	return threads, nil
// }
