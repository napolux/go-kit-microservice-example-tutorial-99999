package napodate

import (
	"context"
	"time"
)

// Service provides some "date capabilities" to your application
type Service interface {
	Status(ctx context.Context) (string, error)
	Get(ctx context.Context) (string, error)
	Validate(ctx context.Context, date string) (bool, error)
}

type dateService struct{}

// NewService makes a new Service.
func NewService() Service {
	return dateService{}
}

// Status only tell us that our service is ok!
func (dateService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

// Get will return today's date
func (dateService) Get(ctx context.Context) (string, error) {
	now := time.Now()
	return now.Format("02/01/2006"), nil
}

// Validate will check if the date today's date
func (dateService) Validate(ctx context.Context, date string) (bool, error) {
	_, err := time.Parse("02/01/2006", date)
	if err != nil {
		return false, err
	}
	return true, nil
}
