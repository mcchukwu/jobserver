// Package service provides decoupleable job services
package service

import (
	"context"
	"time"

	"github.com/mcchukwu/jobserver/internal/job"
)

type JobService struct{}

func NewJobService() *JobService {
	return &JobService{}
}

func (s *JobService) Run(ctx context.Context, input job.JobInput) (job.JobOutput, error) {
	cfg := job.JobConfig{
		Duration: 30 * time.Second,
		Interval: 3 * time.Second,
	}

	return job.RunJob(ctx, input, cfg)
}
