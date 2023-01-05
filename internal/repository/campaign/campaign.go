package campaign

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/hugo.rojas/custom-api/internal/domain/entities"
	"github.com/hugo.rojas/custom-api/internal/domain/models"
	dberrors "github.com/hugo.rojas/custom-api/internal/errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type repoCampaign struct {
	DB *sqlx.DB
}

func NewCampaignRepository(db *sqlx.DB) CampaignRepository {
	return &repoCampaign{
		DB: db,
	}
}

func (r *repoCampaign) Noop(ca entities.Campaign) error {
	return nil
}

func (r *repoCampaign) GetCampaignByID(request entities.GetCampaignRequest) (models.Campaign, *dberrors.ErrorResponse) {
	args := []interface{}{request.ID}
	q := strings.Builder{}
	q.WriteString(`
		SELECT id,name,created_at,updated_at
		FROM campaigns
		WHERE id = $1
	`)

	var campaign models.Campaign
	rowErr := r.DB.QueryRow(q.String(), args...).Scan(&campaign.ID, &campaign.Name, &campaign.CreatedAt, &campaign.UpdatedAt, &campaign.DeletedAt)
	if rowErr != nil {

		if errors.Is(rowErr, sql.ErrNoRows) {
			return campaign, dberrors.NotFoundError("campaign not found")
		}

		return campaign, dberrors.InternalError(rowErr.Error())
	}

	return campaign, nil
}
