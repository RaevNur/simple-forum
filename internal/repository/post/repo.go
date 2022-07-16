package post

import (
	"database/sql"
	"fmt"
	"forum/internal/helper"
	"forum/internal/helper/constraints"

	model "forum/internal/models"
)

func (p *PostRepo) Create(post *model.Post) error {
	query := `INSERT INTO posts (
		content, 
		user_id, 
		created_at
	) 
	VALUES (?, ?, ?);`

	stmt, err := p.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("PostRepo.Create: %w", err)
	}

	defer stmt.Close()

	encodedTime := helper.EncodeTime((*post).CreatedAt, constraints.TimeFormatRFC3339)
	res, err := stmt.Exec((*post).Content, (*post).UserId, encodedTime)
	if err != nil {
		return fmt.Errorf("PostRepo.Create: %w", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("PostRepo.Create: %w", err)
	}

	(*post).Id = lastId
	return nil
}

func (p *PostRepo) GetById(id int64) (*model.Post, error) {
	query := `SELECT posts.content, posts.user_id, posts.created_at FROM posts 
	WHERE posts.id = ?`
	row := p.db.QueryRow(query, id)

	post := model.Post{
		Id: id,
	}
	var encodedTime string

	err := row.Scan(&post.Content, &post.UserId, &encodedTime)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("PostRepo.GetById: %w", err)
	}

	post.CreatedAt, err = helper.DecodeTime(encodedTime, constraints.TimeFormatRFC3339)
	if err != nil {
		return nil, fmt.Errorf("PostRepo.GetById: %w", err)
	}

	return &post, nil
}

func (p *PostRepo) GetLikesCount(post *model.Post) error {
	query := `SELECT COUNT(likes.id) FROM likes
	WHERE likes.post_id = ? AND likes.liked = ?`
	row := p.db.QueryRow(query, post.Id, constraints.LikeValue)

	err := row.Scan(&post.Likes)
	if err == sql.ErrNoRows {
	} else if err != nil {
		return fmt.Errorf("PostRepo.GetLikesCount: %w", err)
	}

	row = p.db.QueryRow(query, post.Id, constraints.DislikeValue)

	err = row.Scan(&post.Dislikes)
	if err == sql.ErrNoRows {
	} else if err != nil {
		return fmt.Errorf("PostRepo.GetLikesCount: %w", err)
	}

	return nil
}

// returns constraints LikeValue(1) and DislikeValue(-1)
// 0 if not liked/disliked
func (p *PostRepo) IsLikedByUser(userId, postId int64) (int, error) {
	query := `SELECT liked FROM likes WHERE user_id = ? AND post_id = ?`
	row := p.db.QueryRow(query, userId, postId)

	var value int

	err := row.Scan(&value)
	if err == sql.ErrNoRows {
		return 0, nil
	} else if err != nil {
		return 0, fmt.Errorf("PostRepo.IsLikedByUser: %w", err)
	}

	return value, nil
}

// returns posts liked by user
// func (p *PostRepo) GetByLiked(userId int64, page int) ([]*model.Post, error) {
// 	query := `SELECT posts.*, IFNULL(SUM(likes.liked), 0) FROM (SELECT posts.id, posts.content, posts.user.id, posts.created_at FROM posts
// 																	INNER JOIN likes ON likes.post_id = posts.id
// 																	WHERE likes.user_id = ? AND likes.liked = ?) AS posts
// 	RIGHT JOIN likes ON likes.post_id = posts.id
// 	GROUP BY posts.id
// 	ORDER BY rate DESC, posts.created_at ASC
// 	LIMIT ? OFFSET ?`

// 	offset := (page - 1) * constraints.LimitThreadsPerPage
// 	rows, err := p.db.Query(query, userId, constraints.LikeValue, constraints.LimitThreadsPerPage, offset)
// 	if err != nil {
// 		return nil, fmt.Errorf("PostRepo.GetByLiked: %w", err)
// 	}

// 	posts := make([]*model.Post, 0, constraints.LimitThreadsPerPage)
// 	for rows.Next() {
// 		t := model.Post{}
// 		var decodedTime string

// 		err = rows.Scan(
// 			&t.Id,
// 			&t.Content,
// 			&t.UserId,
// 			&decodedTime,
// 			&t.Rate,
// 		)
// 		if err != nil {
// 			return nil, fmt.Errorf("PostRepo.GetByLiked: %w", err)
// 		}

// 		t.CreatedAt, err = helper.DecodeTime(decodedTime, constraints.TimeFormatRFC3339)
// 		if err != nil {
// 			return nil, fmt.Errorf("PostRepo.GetByLiked: %w", err)
// 		}
// 		posts = append(posts, &t)
// 	}

// 	return posts, nil
// }
