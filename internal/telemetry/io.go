package telemetry

import (
	"context"

	"github.com/hugo.rojas/custom-api/internal/domain/entities"
	"github.com/hugo.rojas/custom-api/internal/iface"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type IO struct {
	io   iface.IO
	name string
	tp   trace.TracerProvider
}

func NewIO(io iface.IO, tp trace.TracerProvider) iface.IO {
	return &IO{
		io:   io,
		name: "io",
		tp:   tp,
	}
}

func (i *IO) SaveRoom(ctx context.Context, room *entities.Room) error {
	ctx, span := otel.Tracer(i.name).Start(ctx, "SaveRoom")
	defer span.End()

	return i.io.SaveRoom(ctx, room)
}

func (i *IO) FilterUsers(ctx context.Context) ([]entities.User, error) {
	ctx, span := otel.Tracer(i.name).Start(ctx, "FilterUsers")
	defer span.End()

	return i.io.FilterUsers(ctx)
}

func (i *IO) LoginUser(ctx context.Context, user *entities.User) error {
	ctx, span := otel.Tracer(i.name).Start(ctx, "LoginUser")
	defer span.End()

	return i.io.LoginUser(ctx, user)
}
