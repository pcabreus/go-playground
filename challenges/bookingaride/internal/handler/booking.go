package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pcabreus/go-playground/challenges/bookingaride/internal/domain"
	"github.com/pcabreus/go-playground/challenges/bookingaride/internal/ports"
	"github.com/pcabreus/go-playground/challenges/bookingaride/internal/service"
)

// Request and Response structs
//
//	{
//	  "pickup": {"lat": 37.7749, "lng": -122.4194},
//	  "dropoff": {"lat": 37.7849, "lng": -122.4094},
//	  "date": "2025-12-31",
//	  "time": "14:30"
//	}
type CreateBookingRequest struct {
	Pickup  Location `json:"pickup"`
	Dropoff Location `json:"dropoff"`
	Date    string   `json:"date,omitempty"` // 2025-12-31
	Time    string   `json:"time,omitempty"` // 14:30
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type CreateBookingResponse struct {
	Quote     float64 `json:"quote"`
	BookingID string  `json:"booking_id"`
}

type Handler struct {
	Store ports.BookingStore
}

func (h Handler) CreateBooking(ctx *fiber.Ctx) error {
	// receive pickup, dropoff, date, time
	var req CreateBookingRequest

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON("invalid request body")
	}

	// validate input

	// respond with quote and booking ID

	// create estimated price or quote
	input := service.CreateBookingInput{
		Pickup:  domain.Location{Lat: req.Pickup.Lat, Lng: req.Pickup.Lng},
		Dropoff: domain.Location{Lat: req.Dropoff.Lat, Lng: req.Dropoff.Lng},
		Date:    req.Date,
		Time:    req.Time,
	}

	booking, err := service.CreateBooking(ctx.UserContext(), input, h.Store)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON("failed to create booking")
	}

	response := CreateBookingResponse{
		Quote:     booking.EstimatedPrice,
		BookingID: booking.ID,
	}

	return ctx.JSON(response)

}
