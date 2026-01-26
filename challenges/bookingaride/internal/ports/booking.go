package ports

import (
	"context"

	"github.com/pcabreus/go-playground/challenges/bookingaride/internal/domain"
)

// --- INPUT PORTS (Primary/Driving Adapters) ---

// BookingUseCase defines the business operations for booking management
// This is the INPUT PORT that primary adapters (HTTP, CLI, gRPC) depend on
type BookingUseCase interface {
	CreateBooking(ctx context.Context, input CreateBookingInput) (*domain.Booking, error)
}

// CreateBookingInput represents the data needed to create a booking
type CreateBookingInput struct {
	Pickup  domain.Location
	Dropoff domain.Location
	Date    string
	Time    string
}

// --- OUTPUT PORTS (Secondary/Driven Adapters) ---

// BookingStore is the output port for booking persistence
type BookingStore interface {
	SaveBooking(ctx context.Context, booking domain.Booking) error
	GetBooking(ctx context.Context, id string) (*domain.Booking, error)
}

// PricingService is the output port for pricing calculations
type PricingService interface {
	CalculatePrice(ctx context.Context, pickup, dropoff domain.Location) (float64, error)
}
