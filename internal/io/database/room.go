package database

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hugo.rojas/custom-api/internal/domain/entities"
)

var (
	createRoom = `INSERT into rooms (id,name,created_at,updated_at) VALUES ($1,$2,$3,$4)`
)

func (d *Database) GetCampaign(ctx context.Context, campaign *entities.Room) error {
	args := []interface{}{campaign.ID.String()}
	q := strings.Builder{}
	q.WriteString(`
		SELECT id,name,created_at,updated_at
		FROM campaigns
		WHERE id = $1
	`)

	rowErr := d.DB.QueryRow(q.String(), args...).Scan(&campaign.ID, &campaign.Name, &campaign.CreatedAt, &campaign.UpdatedAt)
	if rowErr != nil {

		if errors.Is(rowErr, sql.ErrNoRows) {
			return errors.New("campaign not found")
		}

		return errors.New("campaign not found")
	}

	return nil
}

func (d *Database) SaveRoom(ctx context.Context, room *entities.Room) error {

	if room.ID == uuid.Nil {
		return d.createRoom(ctx, room)
	}
	panic("Not implemented")
}

func (d *Database) createRoom(ctx context.Context, room *entities.Room) error {
	now := time.Now()
	ID := uuid.New()
	_, err := d.DB.Exec(createRoom,
		ID.String(),
		room.Name,
		now, now,
	)
	if err != nil {
		return err
	}

	if err == nil {
		room.CreatedAt = now
		room.UpdatedAt = now
		room.ID = ID
	}
	return nil
}
