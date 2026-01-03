// Package job provides a cancellable job execution logic
package job

import (
	"context"
	"errors"
	"time"
)

var ErrJobCancelled = errors.New("job was cancelled")

type JobConfig struct {
	Duration time.Duration
	Interval time.Duration
}

type JobInput struct {
	A int
	B int
}

type JobOutput struct {
	Sum int
}

func RunJob(ctx context.Context, input JobInput, cfg JobConfig) (JobOutput, error) {
	deadline := time.Now().Add(cfg.Duration)

	// Simulate work
	for {
		select {
		case <-ctx.Done():
			return JobOutput{}, ErrJobCancelled
		case <-time.After(cfg.Interval):
			// Simulate step
		}

		if time.Now().After(deadline) {
			break
		}
	}

	return JobOutput{Sum: input.A + input.B}, nil
}
