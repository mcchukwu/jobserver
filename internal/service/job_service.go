// Package service provides decoupleable job services
package service

import (
	"context"
	"time"

	"github.com/mcchukwu/jobserver/internal/job"
)

func RunJob(ctx context.Context, input job.JobInput) (job.JobOutput, error) {
	duration := 30 * time.Second
	interval := 3 * time.Second
	return job.RunJob(ctx, input, job.JobConfig{Duration: duration, Interval: interval})
}
