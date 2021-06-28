package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

const linkTable = "link"

type Config struct {
	DriverName     string
	DataSourceName string
	MaxFileSize    int64
}

func NewSqlite3DB(config Config) (*sql.DB, error) {
	// удаляем файл, если он превысил максимальный размер (чтобы тестовый стенд не разрастался)
	if fileExists(config.DataSourceName) && fileOverSized(config.DataSourceName, config.MaxFileSize) {
		err := os.Remove(config.DataSourceName)
		if err != nil {
			return nil, err
		}
	}

	db, err := sql.Open(config.DriverName, config.DataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func fileExists(dataSourceName string) bool {
	info, err := os.Stat(dataSourceName)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func fileOverSized(dataSourceName string, maxFileSize int64) bool {
	info, err := os.Stat(dataSourceName)
	if os.IsNotExist(err) {
		return false
	}
	return info.Size() >= maxFileSize
}
