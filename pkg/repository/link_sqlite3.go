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
	// укорачивание ссылки
	// сохранение в бд

	userLink := shorter.UserLink{}

	tx, err := l.db.Begin()
	if err != nil {
		return userLink, err
	}

	// записываю длин ссыль
	// получаю
	var linkId int64
	createItemQuery := fmt.Sprintf("INSERT INTO %s (short, long) values (NULL, $1) RETURNING id", linkTable)
	row := tx.QueryRow(createItemQuery, longLink)
	err = row.Scan(&linkId)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return userLink, rollbackErr
		}
		return userLink, err
	}

	// генер корот
	shortLink := "https://" + os.Getenv("DOMAIN") + "/" + l.notation.Convert(linkId)
	fmt.Println("!!!shortLinkTail ", shortLink)

	// update
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

// generateShortLink - создение короткой ссылки
// Короткие ссылки должны основываться на id записи(sequence) в БД, переведённой в систему счисления с алфавитом [A-Za-z0-9]
func generateShortLink(recordId int) (string, error) {
	// представим, что есть символьно-цифровой алфафит, представленный с помощью A-Za-z0-9
	// A-1, B-2, ... Z-26, a-27, b-28, ... z-52, 0-53, 1-54, ... 9-62, и далее по-кругу: AA-63
	// получилась 63-ричная система счисления
	// теперь просто переводим наш ID в эту систему счисления, напрмиер 1=>A, ... 64=>AB
	return "", nil
}
