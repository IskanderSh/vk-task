package storage

import (
	"testing"
	"time"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestCreateActor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewActorStorage(db)

	type args struct {
		name     string
		sex      string
		birthday time.Time
	}

	type mockBehavior func(args args)

	testTable := []struct {
		name         string
		args         args
		mockBehavior mockBehavior
		wantErr      bool
		err          error
	}{
		{
			name: "OK",
			args: args{
				name:     "Iskander",
				sex:      "male",
				birthday: time.Date(2004, 10, 11, 0, 0, 0, 0, time.UTC),
			},
			mockBehavior: func(args args) {
				mock.ExpectExec("INSERT INTO actors").WithArgs(args.name, args.sex, args.birthday).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Duplicate name",
			args: args{
				name:     "Duplicated",
				sex:      "female",
				birthday: time.Date(2004, 10, 11, 0, 0, 0, 0, time.UTC),
			},
			mockBehavior: func(args args) {
				mock.ExpectExec("INSERT INTO actors").WithArgs(args.name, args.sex, args.birthday).
					WillReturnError(&pq.Error{Code: "23505"})
			},
			wantErr: true,
			err:     ErrDuplicateName,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args)

			err = storage.CreateActor(testCase.args.name, testCase.args.sex, testCase.args.birthday)
			if testCase.wantErr {
				assert.Error(t, err)
				assert.Equal(t, ErrDuplicateName, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
