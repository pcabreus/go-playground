package ports

import (
	"context"

	"github.com/pcabreus/go-playground/challenges/bookingaride/internal/domain"
)

type BookingStore interface {
	SaveBooking(ctx context.Context, booking domain.Booking) error
}
