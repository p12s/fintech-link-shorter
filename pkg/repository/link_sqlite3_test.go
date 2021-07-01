package repository

import (
	"database/sql"
	shorter "github.com/p12s/fintech-link-shorter"
	//"github.com/p12s/fintech-link-shorter/pkg/repository"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"log"
	"testing"
)

func TestLinkSqlite3_Convert10To62(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf(err.Error())
		}
	}(db)

	r := NewLinkSqlite3(db)

	type args struct {
		id int
		item   shorter.Link
	}
	type mockBehavior func(args args, id int)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				id: 1,
				item: shorter.Link{
					Id: 1,
					Long: "https://www.golangprograms.com/example-to-handle-get-and-post-request-in-golang.html",
					Short: "",
				},
			},
			want: 1,
			mock: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id", "long", "short"}).AddRow(args.item.Id, args.item.Long, args.item.Short)
				mock.ExpectQuery("INSERT INTO link").WithArgs(args.item.Long).WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO link").WithArgs(args.item.Long).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input, tt.want)

			got, err := r.Create(tt.input.item.Long)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
