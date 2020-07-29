package basicgokitfirst

import (
	"context"
	"time"
)

// Service for declare our interface
type Service interface {
	Status(ctx context.Context) (string, error)
	Get(ctx context.Context) (string, error)
	Validate(ctx context.Context, date string) (bool, error)
}

// implementation of interface

// DateService is a struct
type DateService struct{}

// NewService make a new Service
// func NewService() Service {
// 	return dateService{}
// }

// Status is a function
func (DateService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

// Get is a function
func (DateService) Get(ctx context.Context) (string, error) {
	now := time.Now()
	return now.Format("02/01/2006"), nil
}

// Validate is a function
func (DateService) Validate(ctx context.Context, date string) (bool, error) {
	_, err := time.Parse("02/01/2006", date)

	if err != nil {
		return false, err
	}

	return true, nil
}
