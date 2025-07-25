package persistence

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// SQLiteStore: wrapper for SQLite DB
// implements the PStore interface
type SQLiteStore struct {
	db *sql.DB
}

// NewSQLiteStore(): to open/create an SQLite DB at the given path
func NewSQLiteStore(path string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// create a new table if it doesn't already exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS kv (
        key TEXT PRIMARY KEY,
        type INTEGER,
        string TEXT,
        list TEXT,
        expiry INTEGER
    )`)
	if err != nil {
		return nil, err
	}

	return &SQLiteStore{db: db}, nil
}

// Save(): to write a slice of serialized values into the db
// replaces values with new data if key already exists
// overwrites previous data in the db
func (s *SQLiteStore) Save(data []SerializedValue) error {

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	// Wipe out old data
	_, err = tx.Exec("DELETE FROM kv")
	if err != nil {
		tx.Rollback()
		return err
	}

	stmt, err := tx.Prepare("REPLACE INTO kv(key, type, string, list, expiry) VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range data {
		listJSON, _ := json.Marshal(item.List)
		_, err := stmt.Exec(item.Key, item.Type, item.String, string(listJSON), item.Expiry)
		if err != nil {
			log.Println("Error saving key:", item.Key, err)
		}
	}
	return tx.Commit()
}

// Load(): read all data from the db as a slice of serialized values
func (s *SQLiteStore) Load() ([]SerializedValue, error) {
	rows, err := s.db.Query("SELECT key, type, string, list, expiry FROM kv")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []SerializedValue
	for rows.Next() {
		var item SerializedValue
		var listJSON string
		if err := rows.Scan(&item.Key, &item.Type, &item.String, &listJSON, &item.Expiry); err != nil {
			return nil, err
		}
		json.Unmarshal([]byte(listJSON), &item.List)
		data = append(data, item)
	}
	return data, nil
}
