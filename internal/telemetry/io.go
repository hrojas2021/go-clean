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

func (i *IO) GetCampaign(ctx context.Context, campaign *entities.Campaign) error {
	ctx, span := otel.Tracer(i.name).Start(ctx, "GetCampaign")
	defer span.End()

	return i.io.GetCampaign(ctx, campaign)
}
