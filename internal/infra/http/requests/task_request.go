package requests

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type TaskRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Deadline    int64  `json:"deadline" validate:"required"`
}

func (r TaskRequest) ToDomainModel() (interface{}, error) {
	return domain.Task{
		Name:        r.Name,
		Description: r.Description,
		Deadline:    time.Unix(r.Deadline, 0),
	}, nil
}
