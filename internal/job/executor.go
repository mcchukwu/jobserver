package job

import (
	"context"
	"errors"
	"time"
)

var ErrJobCancelled = errors.New("Job Cancelled!")

type JopInput struct {
	num1     int
	num2     int
	Duration time.Duration
	Interval time.Duration
}

type JobResult struct {
	sum int
}

func RunJob(ctx context.Context, input JopInput) (JobResult, error) {
	deadline := time.Now().Add(input.Duration)

	for {
		select {
		case <-ctx.Done():
			return JobResult{}, ErrJobCancelled
		default:
			// continue work
		}

		if time.Now().After(dealine) {
			break
		}

		// simulate work
		time.Sleep(input.Interval)
	}

	return JobResult{
		sum: input.num1 + input.num2
	}, nil
}
