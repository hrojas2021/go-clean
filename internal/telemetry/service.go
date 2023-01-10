package telemetry

import (
	"context"

	"github.com/hugo.rojas/custom-api/internal/domain/entities"
	"github.com/hugo.rojas/custom-api/internal/domain/models"
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

func (s *Service) SaveRoom(ctx context.Context, room models.Room) (entities.Room, error) {
	ctx, span := otel.Tracer(s.name).Start(ctx, "SaveRoom")
	defer span.End()

	return s.service.SaveRoom(ctx, room)
}

func (s *Service) ListUser(ctx context.Context) ([]entities.User, error) {
	ctx, span := otel.Tracer(s.name).Start(ctx, "ListUser")
	defer span.End()

	return s.service.ListUser(ctx)
}

func (s *Service) Login(ctx context.Context, user models.User) error {
	ctx, span := otel.Tracer(s.name).Start(ctx, "LoginUser")
	defer span.End()

	return s.service.Login(ctx, user)
}
