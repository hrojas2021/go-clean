package entities

import (
	"time"

	"github.com/google/uuid"
)

// Room model
type Room struct {
	ID        uuid.UUID ` db:"id"         json:"id"`
	Name      string    ` db:"name"       json:"name"`
	CreatedAt time.Time ` db:"created_at" json:"createdAt"`
	UpdatedAt time.Time ` db:"updated_at" json:"updatedAt"`
}
