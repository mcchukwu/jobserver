// Packege Api defines handler logic for this application
package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mcchukwu/jobserver/internal/job"
	"github.com/mcchukwu/jobserver/internal/service"
)

type JobHandler struct {
	s *service.JobService
}

func NewJobHandler(svc *service.JobService) *JobHandler {
	return &JobHandler{s: svc}
}

func (h JobHandler) CreateJob(w http.ResponseWriter, r http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid http request method", http.StatusMethodNotAllowed)
		return
	}

	var input job.JobInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	output, err := h.s.Run(ctx, input)
	if err != nil {
		if errors.Is(err, job.ErrJobCancelled) {
			http.Error(w, "Cancelled", http.StatusRequestTimeout)
		}

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
