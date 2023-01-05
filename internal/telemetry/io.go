package telemetry

import (
	"github.com/hugo.rojas/custom-api/internal/iface"
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
