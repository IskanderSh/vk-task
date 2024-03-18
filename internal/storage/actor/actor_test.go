package actor

import (
	"testing"
	"time"

	"github.com/IskanderSh/vk-task/internal/entities"
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

	storage := NewStorage(db)

	type mockBehavior func(actor entities.Actor)

	testTable := []struct {
		name         string
		actor        entities.Actor
		mockBehavior mockBehavior
		wantErr      bool
		err          error
	}{
		{
			name: "OK",
			actor: entities.Actor{
				Name:     "Iskander",
				Sex:      "male",
				Birthday: time.Date(2004, 10, 11, 0, 0, 0, 0, time.UTC),
			},
			mockBehavior: func(actor entities.Actor) {
				mock.ExpectExec("INSERT INTO actors").WithArgs(actor.Name, actor.Sex, actor.Birthday).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Duplicate name",
			actor: entities.Actor{
				Name:     "Duplicated",
				Sex:      "female",
				Birthday: time.Date(2004, 10, 11, 0, 0, 0, 0, time.UTC),
			},
			mockBehavior: func(actor entities.Actor) {
				mock.ExpectExec("INSERT INTO actors").WithArgs(actor.Name, actor.Sex, actor.Birthday).
					WillReturnError(&pq.Error{Code: "23505"})
			},
			wantErr: true,
			err:     ErrDuplicateName,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.actor)

			err = storage.CreateActor(&testCase.actor)
			if testCase.wantErr {
				assert.Error(t, err)
				assert.Equal(t, testCase.err, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
