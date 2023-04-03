package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/hugo.rojas/custom-api/internal/domain/entities"
	"github.com/hugo.rojas/custom-api/internal/domain/models"
	"github.com/hugo.rojas/custom-api/internal/iface/mock"
	"github.com/hugo.rojas/custom-api/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateRoom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	id := uuid.New()
	now := time.Now()
	fakeRoom := &models.Room{Name: "Fake Room"}
	errInternalTestError := errors.New("internal test error")

	t.Run("Service success", func(t *testing.T) {
		m := mock.NewMockIO(ctrl)
		srv := service.New(nil, m)

		m.EXPECT().SaveRoom(ctx, gomock.Any()).
			DoAndReturn(func(_ context.Context, r *entities.Room) error {
				assert.Equal(t, "Fake Room", r.Name)
				r.ID = id
				r.CreatedAt = now
				r.UpdatedAt = now
				return nil
			})
		assert.Nil(t, srv.SaveRoom(ctx, fakeRoom))
		assert.Equal(t, fakeRoom.ID, id)
		assert.Equal(t, fakeRoom.CreatedAt, now)
		assert.Equal(t, fakeRoom.UpdatedAt, now)
	})

	t.Run("Service error", func(t *testing.T) {
		m := mock.NewMockIO(ctrl)
		m.EXPECT().SaveRoom(gomock.Any(), gomock.Any()).
			Return(errInternalTestError)

		srv := service.New(nil, m)
		err := srv.SaveRoom(ctx, fakeRoom)
		require.Error(t, err)
		assert.Equal(t, "internal test error", err.Error())
	})
}
