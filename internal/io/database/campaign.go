package database

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/hugo.rojas/custom-api/internal/domain/entities"
)

func (d *Database) GetCampaign(ctx context.Context, campaign *entities.Campaign) error {
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
