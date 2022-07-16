package constraints

import "time"

// standard variables
var (
	DB_PATH     = "."
	DB_NAME     = "sqlite.db"
	DB_USERNAME = ""
	DB_PASSWORD = ""
)

const (
	DB_AUTHCRYPT = "SHA256"

	DBDriverName = "sqlite3"

	// db constants
	// TimeFormatRFC1123 = "Mon, 02 Jan 2006 15:04:05 MST"
	TimeFormatRFC3339 = "2006-01-02T15:04:05Z07:00"
	TimeFormatView    = "2006-01-02 15:04:05"
	DislikeValue      = -1
	LikeValue         = 1

	// frontend constants
	LimitThreadsPerPage  = 10
	LimitCommentsPerPage = 10
	LimitTagsPerPage     = 25

	// cookie constants
	CookieExpireTime = 30 * time.Minute
)
