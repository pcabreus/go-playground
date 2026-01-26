package store

import (
	"context"

	"github.com/pcabreus/go-playground/challenges/bookingaride/internal/domain"
)

type BookingStore struct{}

func (b BookingStore) SaveBooking(ctx context.Context, booking domain.Booking) error {

	// out-of-scope: implement actual storage logic (e.g., database)
	return nil
}
