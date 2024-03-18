package handlers

import (
	"context"
	"testing"
	"time"

	"github.com/IskanderSh/vk-task/internal/entities"
	"github.com/IskanderSh/vk-task/internal/generated/models"
	mock_handlers "github.com/IskanderSh/vk-task/internal/services/mocks"
	"github.com/golang/mock/gomock"
)

var (
	name      = "Iskander"
	maleSex   = "male"
	femaleSex = "female"
)

func TestCreateActor(t *testing.T) {
	type mockBehavior func(r *mock_handlers.MockActorProvider, actor models.Actor)

	tests := []struct {
		name                 string
		inputBody            string
		inputActor           models.Actor
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name": "Iskander", "sex": "male", "birthday": "2004-11-10"}`,
			inputActor: entities.ConvertToInput(entities.Actor{
				Name:     "Iskander",
				Sex:      "male",
				Birthday: time.Time{},
			}),
			mockBehavior: func(r *mock_handlers.MockActorProvider, actor models.Actor) {
				r.EXPECT().AddActor(context.Background(), actor).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `successfully created new actor`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//log := &slog.Logger{}

			c := gomock.NewController(t)
			defer c.Finish()

			//repo := services.NewActorService(log, c)
		})
	}
}
