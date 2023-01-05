package telemetry

import (
	"github.com/hugo.rojas/custom-api/internal/iface"
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
