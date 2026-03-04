package db

import (
	"database/sql"
	"path/filepath"

	// SQLite driver
	_ "github.com/mattn/go-sqlite3"
)

// InitDB sets up the SQLite database connection and runs migrations.
// Returns the connection pool or an error if initialization fails.
func InitDB(dataDir string) (*sql.DB, error) {
	dbPath := filepath.Join(dataDir, "gh-stats.sqlite")
	
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := runMigrations(db); err != nil {
		return nil, err
	}

	return db, nil
}
