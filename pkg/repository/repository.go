package repository

import (
	"database/sql"
	"github.com/p12s/fintech-link-shorter"
)

// Link - последний (третий) уровень "луковой архитектуры"
// выполняет работу с БД
type Link interface {
	Create(longLink string) (shorter.UserLink, error)
	Long(shortLink string) (shorter.UserLink, error)
}

// Repository - репозиторий
type Repository struct {
	Link
}

// NewRepository - конструктор
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Link: NewLinkSqlite3(db),
	}
}
