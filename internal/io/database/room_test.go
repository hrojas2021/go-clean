package database

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hugo.rojas/custom-api/internal/domain/entities"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateRoom(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mdb, mock, err := sqlmock.New()
	require.Nil(t, err)
	defer mdb.Close()

	now := time.Now()
	room := entities.Room{Name: "Fake Room"}
	errInternalTestError := errors.New("internal query error")
	t.Run("IO success", func(t *testing.T) {
		t.Parallel()
		sqlxDB := sqlx.NewDb(mdb, "postgres")
		io := New(sqlxDB)

		mock.
			ExpectExec(
				regexp.QuoteMeta(createRoom),
			).WithArgs(sqlmock.AnyArg(), room.Name, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		require.Nil(t, io.SaveRoom(ctx, &room))
		require.True(t, room.CreatedAt.After(now))
		require.True(t, room.UpdatedAt.After(now))
	})

	t.Run("Query error - Io Fail", func(t *testing.T) {
		t.Parallel()
		sqlxDB := sqlx.NewDb(mdb, "postgres")
		io := New(sqlxDB)

		mock.
			ExpectExec(
				regexp.QuoteMeta(createRoom),
			).WithArgs(sqlmock.AnyArg(), "Room", sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(errInternalTestError)

		err := io.SaveRoom(ctx, &entities.Room{Name: "Room"})
		require.Error(t, err)
		assert.Equal(t, "internal query error", err.Error())
	})

	t.Run("Rows affected - Io Fail", func(t *testing.T) {
		t.Parallel()
		sqlxDB := sqlx.NewDb(mdb, "postgres")
		io := New(sqlxDB)

		mockRows := &MockResult{
			err: errors.New("rows affected error"),
		}

		mock.
			ExpectExec(
				regexp.QuoteMeta(createRoom),
			).WithArgs(sqlmock.AnyArg(), "Room", sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(mockRows)

		err := io.SaveRoom(ctx, &entities.Room{Name: "Room"})
		require.Error(t, err)
		assert.Equal(t, "rows affected error", err.Error())
	})

	t.Run("Rows 0 - Io Fail", func(t *testing.T) {
		t.Parallel()
		sqlxDB := sqlx.NewDb(mdb, "postgres")
		io := New(sqlxDB)

		mock.
			ExpectExec(
				regexp.QuoteMeta(createRoom),
			).WithArgs(sqlmock.AnyArg(), "Room", sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := io.SaveRoom(ctx, &entities.Room{Name: "Room"})
		require.Error(t, err)
		assert.Equal(t, "could not create room; 0 rows affected", err.Error())
	})

	t.Cleanup(func() {
		// Nothing to clean up
	})
}

type MockResult struct {
	lastInsertID, rowsAffected int64
	err                        error
}

func (r *MockResult) LastInsertId() (int64, error) {
	return r.lastInsertID, r.err
}

func (r *MockResult) RowsAffected() (int64, error) {
	return r.rowsAffected, r.err
}
