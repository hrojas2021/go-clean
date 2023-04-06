package database

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/hugo.rojas/custom-api/internal/domain/entities"
)

var (
	createRoom = `INSERT into rooms (id,name,created_at,updated_at) VALUES ($1,$2,$3,$4);`
)

func (d *Database) SaveRoom(ctx context.Context, room *entities.Room) error {
	if room.ID == uuid.Nil {
		return d.createRoom(ctx, room)
	}
	panic("Not implemented")
}

func (d *Database) createRoom(_ context.Context, room *entities.Room) error {
	now := time.Now()
	ID := uuid.New()
	res, err := d.DB.Exec(createRoom,
		ID.String(),
		room.Name,
		now, now,
	)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("could not create room; 0 rows affected")
	}

	room.CreatedAt = now
	room.UpdatedAt = now
	room.ID = ID

	return nil
}
