package repository

import (
	"github.com/magiconair/properties/assert"
	"log"
	"os"
	"testing"
)

func TestNotation_FileExists(t *testing.T) {
	tests := []struct {
		name       string
		filePath   string
		fileExists bool
	}{
		{
			name:       "Detect an existing file",
			filePath:   "./sqlite3_test.go",
			fileExists: true,
		},
		{
			name:       "Does not detect a non-existent file",
			filePath:   "./no-exists-file.db",
			fileExists: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, FileExists(test.filePath), test.fileExists)
		})
	}
}

func TestNotation_FileOverSized(t *testing.T) {
	const MaxFileSize = 5 // MB
	const FilePath = "./test-file-size.db"

	tests := []struct {
		name           string
		fileSize       int
		filePath       string
		needCreateFile bool
		fileOverSized  bool
	}{
		{
			name:           "Not detect small sized file",
			fileSize:       MaxFileSize * 1,
			needCreateFile: false,
			fileOverSized:  false,
		},
		{
			name:           "Not detect small sized file",
			fileSize:       MaxFileSize*1024*1024 - 1,
			needCreateFile: true,
			fileOverSized:  false,
		},
		{
			name:           "Not detect equal sized file",
			fileSize:       MaxFileSize * 1024 * 1024,
			needCreateFile: true,
			fileOverSized:  false,
		},
		{
			name:           "Detect over sized file",
			fileSize:       MaxFileSize*1024*1024 + 1,
			needCreateFile: true,
			fileOverSized:  true,
		},
		{
			name:           "Detect over sized file",
			fileSize:       MaxFileSize * 1024 * 1024 * 2,
			needCreateFile: true,
			fileOverSized:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.needCreateFile {
				_, err := os.Stat(FilePath)
				if os.IsExist(err) {
					if err := os.Remove(FilePath); err != nil {
						log.Fatal("Can't remove test file " + FilePath)
					}
				}
				fd, err := os.Create(FilePath)
				if err != nil {
					log.Fatal("Failed to create output")
				}
				if _, err = fd.Write(make([]byte, test.fileSize)); err != nil {
					log.Fatal("Write failed")
				}

				err = fd.Close()
				if err != nil {
					log.Fatal("Failed to close file")
				}
			}

			assert.Equal(t, FileOverSized(FilePath, MaxFileSize*1024*1024), test.fileOverSized)

			if test.needCreateFile {
				if err := os.Remove(FilePath); err != nil {
					log.Fatal("Can't remove test file " + FilePath)
				}
			}
		})
	}
}
