package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pcabreus/go-playground/challenges/bookingaride/internal/domain"
	"github.com/pcabreus/go-playground/challenges/bookingaride/internal/ports"
)

// BookingService encapsulates booking business logic and implements BookingUseCase
type BookingService struct {
	storagePort ports.BookingStore
	pricePort   ports.PricingService
}

// NewBookingService creates a new instance of BookingService
func NewBookingService(storage ports.BookingStore, pricing ports.PricingService) *BookingService {
	return &BookingService{
		storagePort: storage,
		pricePort:   pricing,
	}
}

// CreateBooking creates a new booking with estimated price
// Implements the BookingUseCase interface
func (bs *BookingService) CreateBooking(ctx context.Context, input ports.CreateBookingInput) (*domain.Booking, error) {
	// Validate input locations
	if err := input.Pickup.Validate(); err != nil {
		return nil, fmt.Errorf("invalid pickup location: %w", err)
	}
	if err := input.Dropoff.Validate(); err != nil {
		return nil, fmt.Errorf("invalid dropoff location: %w", err)
	}

	// Calculate estimated price using the pricing port
	quote, err := bs.pricePort.CalculatePrice(ctx, input.Pickup, input.Dropoff)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrPriceCalculation, err)
	}

	// Generate unique booking ID
	bookingID := uuid.New().String()
	now := time.Now().Format(time.RFC3339)

	// Create domain booking entity
	booking := domain.Booking{
		ID:             bookingID,
		Pickup:         input.Pickup,
		Dropoff:        input.Dropoff,
		Date:           input.Date,
		Time:           input.Time,
		EstimatedPrice: quote,
		Status:         "pending",
		CreatedAt:      now,
		UpdatedAt:      now,
		GuestID:        nil,
	}

	// Validate the complete booking
	if err := booking.Validate(); err != nil {
		return nil, fmt.Errorf("invalid booking: %w", err)
	}

	// Persist the booking using storage port
	if err := bs.storagePort.SaveBooking(ctx, booking); err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrStorageUnavailable, err)
	}

	return &booking, nil
}
