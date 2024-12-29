package db

import (
	"database/sql"
)

// store provide all func to execute the query
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}

}
