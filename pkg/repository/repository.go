package repository

import (
	"database/sql"
	"github.com/p12s/fintech-link-shorter"
)

type Link interface {
	Create(longLink string) (shorter.UserLink, error)
	Long(shortLink string) (shorter.UserLink, error)
}

type Repository struct {
	Link
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Link: NewLinkSqlite3(db),
	}
}
