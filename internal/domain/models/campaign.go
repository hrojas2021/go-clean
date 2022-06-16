package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	null "gopkg.in/guregu/null.v3"
)

// Campaign model
type Campaign struct {
	ID        uuid.UUID ` db:"id"         json:"id"`
	Name      string    ` db:"name"       json:"name"`
	CreatedAt time.Time ` db:"created_at" json:"createdAt"`
	UpdatedAt time.Time ` db:"updated_at" json:"updatedAt"`
	DeletedAt null.Time ` db:"deleted_at" json:"deletedAt"`
}

type GetCampaignRequest struct {
	ID string `json:"id"`
}
