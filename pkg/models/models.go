package models

import (
	"database/sql"
)

// Item represent the item model
type Item struct {
	ID    int64          `db:"id"`
	URL   string         `db:"url"`
	Short sql.NullString `db:"short"`
}

// ResponseWrapper is the format of a wrapper for API Response
type ResponseWrapper struct {
	Result interface{}
	Status int
	Error  string
	Time   string
}
