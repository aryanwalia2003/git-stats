package db

import "database/sql"

const schema = `
-- Table 1: Snapshots of GitHub API statistics (Stars, Forks, Issues)
CREATE TABLE IF NOT EXISTS repo_snapshots (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	repo_origin_url TEXT NOT NULL,
	timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
	stars_count INTEGER NOT NULL,
	forks_count INTEGER NOT NULL,
	open_issues_count INTEGER NOT NULL
);

-- Table 2: Cache for AI Generated Insights to avoid re-prompting LLM
CREATE TABLE IF NOT EXISTS ai_insights (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	repo_origin_url TEXT NOT NULL,
	insight_title TEXT NOT NULL,
	insight_description TEXT NOT NULL,
	insight_type TEXT NOT NULL, -- 'milestone', 'vibe_check', 'refactor'
	generated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`

func runMigrations(db *sql.DB) error {
	_, err := db.Exec(schema)
	return err
}
