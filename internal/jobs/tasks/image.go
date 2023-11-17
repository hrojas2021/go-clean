package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hibiken/asynq"
	"github.com/hugo.rojas/custom-api/internal/iface"
)

type ImageProcessor struct{}

func NewImageProcessor(service iface.Service) *ImageProcessor {
	return &ImageProcessor{}
}

type ImageResizePayload struct {
	SourceURL string
}

func NewImageResizeTask() *asynq.Task {
	src := "Test Task"
	payload, err := json.Marshal(ImageResizePayload{SourceURL: src})
	if err != nil {
		os.Exit(1)
	}

	//task options can be passed to NewTask, which can be overridden at enqueue time
	return asynq.NewTask(TypeImageResize, payload, asynq.MaxRetry(5), asynq.Timeout(5*time.Second))
}

func (ip *ImageProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p ImageResizePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed %v> %w", err, asynq.SkipRetry)
	}
	log.Printf("Resizing image: src=%s", p.SourceURL)
	return nil
}
