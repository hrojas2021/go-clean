package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/hugo.rojas/custom-api/conf"
	"github.com/hugo.rojas/custom-api/internal/domain/models"
	"github.com/hugo.rojas/custom-api/internal/http/rest"
	"github.com/hugo.rojas/custom-api/internal/http/rest/handlers"
	"github.com/hugo.rojas/custom-api/internal/iface/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateRoom(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := uuid.New()
	now := time.Now()
	fakeRoom := &models.Room{Name: "Fake Room"}
	errInternalTestError := errors.New("internal test error")

	t.Run("Success request", func(t *testing.T) {
		// initialize
		m := mock.NewMockService(ctrl)

		// expect mock call
		m.EXPECT().SaveRoom(gomock.Any(), fakeRoom).
			DoAndReturn(func(_ context.Context, r *models.Room) error {
				assert.Equal(t, "Fake Room", r.Name)
				r.ID = id
				r.CreatedAt = now
				r.UpdatedAt = now
				return nil
			})

		// config router
		r := rest.InitRoutes(m, &conf.Configuration{})
		h := handlers.New(m, new(rest.DefaultResp))
		r.POST("/", h.SaveRoom)

		// test server
		ts := httptest.NewServer(r)
		defer ts.Close()

		// prepare and call
		body := bytes.NewBufferString("{\"name\":\"Fake Room\"}")
		res, err := http.Post(fmt.Sprintf("%s/", ts.URL), "application/json", body)

		// validate
		require.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var room models.Room
		require.Nil(t, json.NewDecoder(res.Body).Decode(&room))
		assert.Equal(t, id.String(), room.ID.String())

		res.Body.Close()
	})

	t.Run("Payload error", func(t *testing.T) {
		m := mock.NewMockService(ctrl)

		r := rest.InitRoutes(m, &conf.Configuration{})
		h := handlers.New(m, new(rest.DefaultResp))
		r.POST("/", h.SaveRoom)

		ts := httptest.NewServer(r)
		defer ts.Close()

		body := bytes.NewBufferString("{\"invalidField\"")
		res, err := http.Post(fmt.Sprintf("%s/", ts.URL), "application/json", body)

		require.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)

		var resp struct {
			rest.ErrResponse
			Room models.Room
		}
		require.Nil(t, json.NewDecoder(res.Body).Decode(&resp))
		res.Body.Close()

		assert.Empty(t, resp.Room)
		assert.Equal(t, "could not parse body params; invalid payload; bad request", resp.Error.Msg)
	})

	t.Run("Invalid payload", func(t *testing.T) {
		m := mock.NewMockService(ctrl)

		r := rest.InitRoutes(m, &conf.Configuration{})
		h := handlers.New(m, new(rest.DefaultResp))
		r.POST("/", h.SaveRoom)

		ts := httptest.NewServer(r)
		defer ts.Close()

		body := bytes.NewBufferString("{\"invalidField\":\"This payload is invalid\"}")
		res, err := http.Post(fmt.Sprintf("%s/", ts.URL), "application/json", body)

		require.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		var resp struct {
			rest.ErrResponse
			Room models.Room
		}
		require.Nil(t, json.NewDecoder(res.Body).Decode(&resp))
		res.Body.Close()

		assert.Empty(t, resp.Room)
		assert.Equal(t, "the params are invalid; invalid payload; bad request", resp.Error.Msg)
	})

	t.Run("Service error", func(t *testing.T) {
		m := mock.NewMockService(ctrl)

		m.EXPECT().SaveRoom(gomock.Any(), gomock.Any()).
			Return(errInternalTestError)

		r := rest.InitRoutes(m, nil)
		h := handlers.New(m, new(rest.DefaultResp))
		r.POST("/", h.SaveRoom)

		ts := httptest.NewServer(r)
		defer ts.Close()

		body := bytes.NewBufferString("{\"name\":\"Fake Room\"}")
		res, err := http.Post(fmt.Sprintf("%s/", ts.URL), "application/json", body)

		require.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

		var resp struct {
			rest.ErrResponse
			Room models.Room
		}
		require.Nil(t, json.NewDecoder(res.Body).Decode(&resp))
		res.Body.Close()

		assert.Empty(t, resp.Room)
		assert.Equal(t, "could not save the room; internal test error", resp.Error.Msg)
	})
}
