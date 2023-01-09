package telemetry

import (
	"context"

	"github.com/google/uuid"
	"github.com/hugo.rojas/custom-api/internal/domain/entities"
	"github.com/hugo.rojas/custom-api/internal/iface"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	service iface.Service
	name    string
	tp      trace.TracerProvider
}

func NewService(service iface.Service, tp trace.TracerProvider) iface.Service {
	return &Service{
		service: service,
		name:    "service",
		tp:      tp,
	}
}

func (s *Service) GetCampaign(ctx context.Context, campaignID uuid.UUID) (*entities.Campaign, error) {
	ctx, span := otel.Tracer(s.name).Start(ctx, "GetCampaign")
	defer span.End()

	return s.service.GetCampaign(ctx, campaignID)
}
