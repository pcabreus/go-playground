package main

import (
	"fmt"
	"log"
	"os"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/pcabreus/go-playground/challenges/bookingaride/internal/handler"
	"github.com/pcabreus/go-playground/challenges/bookingaride/internal/store"
)

const (
	DefaultPort = "8080"
)

func main() {
	// require config
	post := os.Getenv("APP_API_PORT")
	if post == "" {
		post = DefaultPort
	}

	addr := fmt.Sprintf(":%s", post)

	app := fiber.New()

	store := store.BookingStore{}
	h := handler.Handler{
		Store: store,
	}

	// routes
	app.Get("/", h.CreateBooking)

	// request quote

	// POST /bookings -> {quote: number, booking_id: string}
	// {pickup: {lat, lng}, dropoff: {lat, lng}, date: string?, time: string?}

	// POST /bookings/{booking_id}/confirm -> {Booking Confirmation with booking_id and estimated_price}
	// {pickup: {lat, lng}, dropoff: {lat, lng}, date: string, time: string, quote: number}

	// GET /bookings -> {Booking Confirmation with booking_id and estimated_price}

	log.Fatal(app.Listen(addr))
}

// Ride
// - Pickup Location (latitude, longitude)
// - Dropoff Location (latitude, longitude)
// - Date and Time of Ride
// - Estimated Price
// - ID
// - Status (e.g., "booked", "completed", "canceled")
// - CreatedAt
// - UpdatedAt
// - GuestId

// Estimation Request
// Receive the pickup and dropoff locations, along with optional date and time.
// Calculate the estimated price based on distance and time factors.
// Use a third-party mapping API to get distance and estimated travel time.
// Save the estimation request to the database.
// Return the estimated price to the user.
