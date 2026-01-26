package service

import (
	"math"

	"github.com/pcabreus/go-playground/challenges/bookingaride/internal/domain"
)

const (
	BaseFare    = 5.00 // tarifa base en dólares
	PerMileRate = 2.00 // tarifa por milla en dólares
)

func CalculateEstimatedPrice(pickup, dropoff domain.Location) float64 {
	distance := math.Sqrt((pickup.Lat-dropoff.Lat)*(pickup.Lat-dropoff.Lat) + (pickup.Lng-dropoff.Lng)*(pickup.Lng-dropoff.Lng))

	// estimatedPrice := baseFare + (perMileRate * distance)
	return distance*PerMileRate + BaseFare
}
