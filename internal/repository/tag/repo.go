package tag

import (
	"database/sql"
	"fmt"

	model "forum/internal/models"
)

func (t *TagRepo) Create(tag *model.Tag) error {
	query := `INSERT INTO tags (
		name
	) 
	VALUES (?);`

	stmt, err := t.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("TagRepo.Create: %w", err)
	}

	defer stmt.Close()

	res, err := stmt.Exec((*tag).Name)
	if err != nil {
		return fmt.Errorf("TagRepo.Create: %w", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("TagRepo.Create: %w", err)
	}

	(*tag).Id = lastId
	return nil
}

// creates realtion in tags_threads table
func (t *TagRepo) CreateRelation(threadId int64, tag *model.Tag) error {
	query := `INSERT INTO tags_threads (
		tag_id,
		thread_id
	) 
	VALUES (?, ?);`

	stmt, err := t.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("TagRepo.CreateRelation: %w", err)
	}

	defer stmt.Close()

	res, err := stmt.Exec((*tag).Id, threadId)
	if err != nil {
		return fmt.Errorf("TagRepo.CreateRelation: %w", err)
	}

	_, err = res.LastInsertId()
	if err != nil {
		return fmt.Errorf("TagRepo.CreateRelation: %w", err)
	}

	return nil
}

// gets all tags ordered by it's amount(desc) and name(asc)
func (t *TagRepo) GetTags() ([]*model.Tag, error) {
	query := `SELECT tags.id, tags.name, COUNT(tg.thread_id) AS "amount" FROM tags 
	INNER JOIN tags_threads AS tg ON tags.id = tg.tag_id 
	GROUP BY tags.id
	ORDER BY amount DESC, tags.name ASC `

	rows, err := t.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("TagRepo.GetTags: %w", err)
	}

	tags := make([]*model.Tag, 0)
	for rows.Next() {
		t := model.Tag{}

		err = rows.Scan(
			&t.Id,
			&t.Name,
			&t.Count,
		)
		if err != nil {
			return nil, fmt.Errorf("TagRepo.GetTags: %w", err)
		}

		tags = append(tags, &t)
	}

	return tags, nil
}

func (t *TagRepo) GetTagByName(name string) (*model.Tag, error) {
	query := `SELECT tags.id, tags.name FROM tags 
	WHERE tags.name = ? `
	row := t.db.QueryRow(query, name)

	tag := model.Tag{}

	err := row.Scan(&tag.Id, &tag.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("TagRepo.GetTagByName: %w", err)
	}

	return &tag, nil
}

// gets tags by page ordered by it's amount(desc) and name(asc)
// func (t *TagRepo) GetTags(page int) ([]*model.Tag, error) {
// 	query := `SELECT tags.id, tags.name, COUNT(tg.thread_id) AS "amount" FROM tags
// 	INNER JOIN tags_threads AS tg ON tags.id = tg.tag_id
// 	GROUP BY tags.id
// 	ORDER BY amount DESC, tags.name ASC
// 	LIMIT ? OFFSET ?`

// 	offset := (page - 1) * constraints.LimitTagsPerPage
// 	rows, err := t.db.Query(query, constraints.LimitTagsPerPage, offset)
// 	if err != nil {
// 		return nil, fmt.Errorf("TagRepo.GetTags: %w", err)
// 	}

// 	tags := make([]*model.Tag, 0, constraints.LimitTagsPerPage)
// 	for rows.Next() {
// 		t := model.Tag{}

// 		err = rows.Scan(
// 			&t.Id,
// 			&t.Name,
// 		)
// 		if err != nil {
// 			return nil, fmt.Errorf("TagRepo.GetTags: %w", err)
// 		}

// 		tags = append(tags, &t)
// 	}

// 	return tags, nil
// }

func (t *TagRepo) GetTagsByThread(threadId int64) ([]*model.Tag, error) {
	query := `SELECT tags.id, tags.name FROM tags 
	INNER JOIN tags_threads AS tg ON tags.id = tg.tag_id
	WHERE tg.thread_id = ? 
	ORDER BY tags.name ASC`

	rows, err := t.db.Query(query, threadId)
	if err != nil {
		return nil, fmt.Errorf("TagRepo.GetTagsByThread: %w", err)
	}

	tags := make([]*model.Tag, 0)
	for rows.Next() {
		t := model.Tag{}

		err = rows.Scan(
			&t.Id,
			&t.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("TagRepo.GetTagsByThread: %w", err)
		}

		tags = append(tags, &t)
	}

	return tags, nil
}
