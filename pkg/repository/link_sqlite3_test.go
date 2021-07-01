package repository

/*
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
	fmt.Println(r)

	type args struct {
		listId int
		item   shorter.UserLink
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
				listId: 1,
				item: shorter.UserLink{
					Url: "https://www.golangprograms.com/example-to-handle-get-and-post-request-in-golang.html",
				},
			},
			want: 2,
			mock: func(args args, id int) {
				/*mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO todo_items").
					WithArgs(args.item.Title, args.item.Description).WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO lists_items").WithArgs(args.listId, id).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input, tt.want)

			got, err := r.Create(tt.input.listId, tt.input.item)
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
*/
