package repository

import (
	"database/sql"
	"github.com/p12s/fintech-link-shorter"
)

type LinkSqlite3 struct {
	db *sql.DB
}

func NewLinkSqlite3(db *sql.DB) *LinkSqlite3 {
	return &LinkSqlite3{db: db}
}

func (l *LinkSqlite3) Create(longLink string) (shorter.UserLink, error) {
	// укорачивание ссылки
	// сохранение в бд

	return shorter.UserLink{
		Url: "https://this-is-short",
	}, nil
}
