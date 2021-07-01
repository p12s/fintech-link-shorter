package repository

import (
	"database/sql"
	"fmt"
	"github.com/p12s/fintech-link-shorter"
	"github.com/p12s/fintech-link-shorter/internal/notation"
	"os"
)

type LinkSqlite3 struct {
	db       *sql.DB
	notation *notation.Convert
}

func NewLinkSqlite3(db *sql.DB) *LinkSqlite3 {
	return &LinkSqlite3{
		db:       db,
		notation: notation.NewConvert(),
	}
}

func (l *LinkSqlite3) Create(longLink string) (shorter.UserLink, error) {
	userLink := shorter.UserLink{}

	tx, err := l.db.Begin()
	if err != nil {
		return userLink, err
	}

	var linkId int64
	createItemQuery := fmt.Sprintf("INSERT INTO %s long values ($1) RETURNING id", linkTable)
	row := tx.QueryRow(createItemQuery, longLink)
	err = row.Scan(&linkId)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return userLink, rollbackErr
		}
		return userLink, err
	}

	shortLink := "https://" + os.Getenv("DOMAIN") + "/" + l.notation.Convert(linkId)
	updateLinkQuery := fmt.Sprintf("UPDATE %s SET short = $1 WHERE id = $2", linkTable)
	_, err = tx.Exec(updateLinkQuery, shortLink, linkId)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return userLink, rollbackErr
		}
		return userLink, err
	}

	return shorter.UserLink{
		Url: shortLink,
	}, tx.Commit()
}

func (l *LinkSqlite3) Long(shortLink string) (shorter.UserLink, error) {
	var link = shorter.Link{}

	query := fmt.Sprintf(`SELECT id, long, short FROM %s WHERE short=$1`, linkTable)
	err := l.db.QueryRow(query, shortLink).Scan(&link.Id, &link.Long, &link.Short)
	if err != nil {
		return shorter.UserLink{}, err
	}

	return shorter.UserLink{Url: link.Long}, nil
}
