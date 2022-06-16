package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
	null "gopkg.in/guregu/null.v3"
)

type Campaign struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type GetCampaignRequest struct {
	ID string
}
