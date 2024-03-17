package user

import (
	"testing"

	"github.com/IskanderSh/vk-task/internal/entities"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	storage := NewStorage(db)

	type mockBehavior func(user entities.CreateUser)

	testTable := []struct {
		name         string
		user         entities.CreateUser
		mockBehavior mockBehavior
		wantErr      bool
		err          error
	}{
		{
			name: "OK",
			user: entities.CreateUser{
				Email:    "admin@admin.com",
				Password: "admin",
				Role:     "admin",
			},
			mockBehavior: func(user entities.CreateUser) {
				mock.ExpectExec("INSERT INTO users").WithArgs(user.Email, user.Password, user.Role).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Duplicate email",
			user: entities.CreateUser{
				Email:    "duplicat@gmail.com",
				Password: "duplicate",
				Role:     "user",
			},
			mockBehavior: func(user entities.CreateUser) {
				mock.ExpectExec("INSERT INTO users").WithArgs(user.Email, user.Password, user.Role).
					WillReturnError(&pq.Error{Code: "23505"})
			},
			wantErr: true,
			err:     ErrDuplicateEmail,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.user)

			err = storage.CreateUser(&testCase.user)
			if testCase.wantErr {
				assert.Error(t, err)
				assert.Equal(t, testCase.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
