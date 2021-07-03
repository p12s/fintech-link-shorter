package repository

import (
	"database/sql"
	"fmt"
	"github.com/p12s/fintech-link-shorter"
	"os"
)

// LinkSqlite3 - структура, содержащая ссылку на БД и конвертер системы счисления
// конвертер используется при каждой записи длинной ссылки в таблицу для декодирования ее в короткую
type LinkSqlite3 struct {
	db       *sql.DB
	notation *Convert
}

// NewLinkSqlite3 - конструктор структуры БД
func NewLinkSqlite3(db *sql.DB) *LinkSqlite3 {
	return &LinkSqlite3{
		db:       db,
		notation: NewConvert(),
	}
}

// Create - запись длинной ссылки в БД с ее конвертацией в короткую
func (l *LinkSqlite3) Create(longLink string) (shorter.UserLink, error) {
	userLink := shorter.UserLink{}

	tx, err := l.db.Begin()
	if err != nil {
		return userLink, err
	}

	var linkId int64
	createItemQuery := fmt.Sprintf("INSERT INTO %s (long) values ($1) RETURNING id", linkTable)
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

// Long - получение длинной ссылки по ее короткой версии
func (l *LinkSqlite3) Long(shortLink string) (shorter.UserLink, error) {
	var link = shorter.Link{}

	query := fmt.Sprintf(`SELECT id, long, short FROM %s WHERE short=$1`, linkTable)
	err := l.db.QueryRow(query, shortLink).Scan(&link.Id, &link.Long, &link.Short)
	if err != nil {
		return shorter.UserLink{}, err
	}

	return shorter.UserLink{Url: link.Long}, nil
}
