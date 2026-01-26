package domain

import (
	"fmt"
)

type Booking struct {
	ID             string
	Pickup         Location
	Dropoff        Location
	Date           string
	Time           string
	EstimatedPrice float64
	Status         string // "pending", "confirmed", "completed", "canceled"
	CreatedAt      string
	UpdatedAt      string
	GuestID        *string
}

type Location struct {
	Lat float64
	Lng float64
}

// Validate checks if the location has valid coordinates
func (l Location) Validate() error {
	if l.Lat < -90 || l.Lat > 90 {
		return fmt.Errorf("%w: latitude must be between -90 and 90, got %f", ErrInvalidLocation, l.Lat)
	}
	if l.Lng < -180 || l.Lng > 180 {
		return fmt.Errorf("%w: longitude must be between -180 and 180, got %f", ErrInvalidLocation, l.Lng)
	}
	return nil
}

// Validate checks if the booking has valid data
func (b Booking) Validate() error {
	if err := b.Pickup.Validate(); err != nil {
		return fmt.Errorf("pickup location: %w", err)
	}
	if err := b.Dropoff.Validate(); err != nil {
		return fmt.Errorf("dropoff location: %w", err)
	}
	if b.EstimatedPrice < 0 {
		return fmt.Errorf("estimated price cannot be negative: %f", b.EstimatedPrice)
	}
	return nil
}
