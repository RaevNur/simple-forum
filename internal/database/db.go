package database

import (
	"database/sql"
	"fmt"
	"forum/internal/helper/constraints"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// initialize connection with sqlite db
// returns (connection, error)
func InitDB() (*sql.DB, error) {
	err := os.MkdirAll(constraints.DB_PATH, os.ModeDir)
	if err != nil {
		return nil, fmt.Errorf("InitDB: %w", err)
	}

	dbPathField := filepath.Join(constraints.DB_PATH, constraints.DB_NAME)

	// connection field from configs
	if constraints.DB_USERNAME != "" && constraints.DB_PASSWORD != "" {
		authField := fmt.Sprintf("?_auth&_auth_user=%s&_auth_pass=%s&_auth_crypt=%s", constraints.DB_USERNAME, constraints.DB_PASSWORD, constraints.DB_AUTHCRYPT)
		dbPathField += authField
	}

	db, err := sql.Open(constraints.DBDriverName, dbPathField)
	if err != nil {
		return nil, fmt.Errorf("InitDB: %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("InitDB: %w", err)
	}

	err = checkDB(db)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("InitDB: %w", err)
	}

	db.SetMaxIdleConns(100)
	return db, err
}

// check the scheme
func checkDB(db *sql.DB) error {
	_, err := db.Exec(constraints.CreateUserTable)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}

	_, err = db.Exec(constraints.CreateSessionsTable)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}

	_, err = db.Exec(constraints.CreateTagsTable)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}

	_, err = db.Exec(constraints.CreatePostsTable)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}

	_, err = db.Exec(constraints.CreateCommentsTable)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}

	_, err = db.Exec(constraints.CreateThreadsTable)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}

	_, err = db.Exec(constraints.CreateLikesTable)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}

	_, err = db.Exec(constraints.CreateTagsThreadsTable)
	if err != nil {
		return fmt.Errorf("checkDB: %w", err)
	}

	return nil
}
