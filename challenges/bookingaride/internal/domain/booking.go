package domain

type Booking struct {
	ID             string
	Pickup         Location
	Dropoff        Location
	Date           string
	Time           string
	EstimatedPrice float64
	Status         string // "booked", "completed", "canceled"
	CreatedAt      string
	UpdatedAt      string
	GuestID        *string
}

type Location struct {
	Lat float64
	Lng float64
}
