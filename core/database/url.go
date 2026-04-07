package database

import "time"

type URL struct {
	ID        uint      `database:"id"`
	URL       string    `database:"url"`
	LongURL   string    `database:"long_url"`
	CreatedAt time.Time `database:"created_at"`
	UpdatedAt time.Time `database:"updated_at"`
	DeletedAt time.Time `database:"deleted_at"`
}
