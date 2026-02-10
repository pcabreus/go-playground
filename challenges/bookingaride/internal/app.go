package internal

import (
	"net/http"
	"time"

	"github.com/pcabreus/challenges/bookingaride/internal/handler"
)

const DateTimeFormat = time.RFC3339 // "2006-01-02T15:04:05Z07:00"

type App struct {
	httpServer *http.ServeMux
	addr       string
}

func New(httpServer *http.ServeMux, addr string) *App {
	return &App{
		httpServer: httpServer,
		addr:       addr,
	}
}

func (a *App) Run() error {

	a.routes()

	return http.ListenAndServe(a.addr, a.httpServer)
}

func (a *App) routes() {

	bookingHandler := handler.Booking{}

	// Add an endpoint to POST a booking request
	a.httpServer.HandleFunc("/bookings", bookingHandler.CreateBooking)

	// add an endpoint to get the booking for a given UUID
	a.httpServer.HandleFunc("/bookings/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		id := r.URL.Path[len("/bookings/"):]
		if id == "" {
			http.Error(w, "Booking ID is required", http.StatusBadRequest)
			return
		}
		w.Write([]byte("Get booking endpoint"))
	})

	a.httpServer.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Booking a Ride Service is running"))
	})

}
