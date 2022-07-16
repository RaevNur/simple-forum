package constraints

const (
	CreateUserTable = `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		nickname TEXT,
		email TEXT,
		created_at TEXT,
		password TEXT,
		first_name TEXT,
		last_name TEXT
	);`
	CreateSessionsTable = `CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY,
		uuid TEXT,
		user_id INTEGER,
		created_at TEXT
	);`
	CreateTagsTable = `CREATE TABLE IF NOT EXISTS tags (
		id INTEGER PRIMARY KEY,
		name TEXT
	);`
	CreatePostsTable = `CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY,
		content TEXT,
		user_id INTEGER,
		created_at TEXT
	);`
	CreateCommentsTable = `CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY,
		post_id INTEGER,
		thread_id INTEGER
	);`
	CreateThreadsTable = `CREATE TABLE IF NOT EXISTS threads (
		id INTEGER PRIMARY KEY,
		post_id INTEGER,
		title TEXT
	);`
	CreateLikesTable = `CREATE TABLE IF NOT EXISTS likes (
		id INTEGER PRIMARY KEY,
		user_id TEXT,
		post_id TEXT,
		liked INTEGER
	);`
	CreateTagsThreadsTable = `CREATE TABLE IF NOT EXISTS tags_threads (
		id INTEGER PRIMARY KEY,
		tag_id INTEGER,
		thread_id INTEGER
	);`
)
