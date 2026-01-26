package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/pcabreus/go-playground/challenges/bookingaride/internal/domain"
	"github.com/pcabreus/go-playground/challenges/bookingaride/internal/ports"
)

type CreateBookingInput struct {
	Pickup  domain.Location
	Dropoff domain.Location
	Date    string
	Time    string
}

func CreateBooking(ctx context.Context, input CreateBookingInput, store ports.BookingStore) (domain.Booking, error) {

	// create estimated price or quote
	quote := CalculateEstimatedPrice(input.Pickup, input.Dropoff)

	UUID := uuid.New()

	BookingID := UUID.String()
	booking := domain.Booking{
		ID:             BookingID,
		Pickup:         domain.Location(input.Pickup),
		Dropoff:        domain.Location(input.Dropoff),
		Date:           input.Date,
		Time:           input.Time,
		EstimatedPrice: quote,
		Status:         "pending",
		GuestID:        nil, // placeholder
	}

	err := store.SaveBooking(ctx, booking)
	if err != nil {
		return domain.Booking{}, err
	}

	return booking, nil
}
