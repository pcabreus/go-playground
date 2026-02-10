package internal

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// func TestCreateBooking(t *testing.T) {
// 	// Arrange
// 	mux := http.NewServeMux()
// 	app := internal.New(mux, ":8080")
// 	app.Run() // This sets up the routes

// 	requestBody := map[string]interface{}{
// 		"guest_id": "g_123",
// 		"pickup": map[string]string{
// 			"query": "Barcelona El Prat T1",
// 		},
// 		"dropoff": map[string]string{
// 			"query": "Passeig de Gràcia 10, Barcelona",
// 		},
// 		"pickup_at": "2026-02-01T10:00:00Z00:00",
// 	}

// 	bodyBytes, err := json.Marshal(requestBody)
// 	if err != nil {
// 		t.Fatalf("Failed to marshal request body: %v", err)
// 	}

// 	req := httptest.NewRequest(http.MethodPost, "/bookings", bytes.NewBuffer(bodyBytes))
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()

// 	// Act
// 	mux.ServeHTTP(w, req)

// 	// Assert
// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusCreated {
// 		t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Fatalf("Failed to read response body: %v", err)
// 	}

// 	var response map[string]interface{}
// 	err = json.Unmarshal(body, &response)
// 	if err != nil {
// 		t.Fatalf("Failed to unmarshal response body: %v", err)
// 	}

// 	// Verify booking_id is a UUID (non-empty string)
// 	bookingID, ok := response["booking_id"].(string)
// 	if !ok || bookingID == "" {
// 		t.Errorf("Expected booking_id to be a non-empty string, got %v", response["booking_id"])
// 	}

// 	// Verify estimation is 99.00
// 	estimation, ok := response["estimation"].(float64)
// 	if !ok {
// 		t.Errorf("Expected estimation to be a number, got %v", response["estimation"])
// 	}
// 	if estimation != 99.00 {
// 		t.Errorf("Expected estimation to be 99.00, got %f", estimation)
// 	}
// }

func TestCreateBooking(t *testing.T) {

	// set dependencies

	httpServer := http.NewServeMux()

	// create app
	app := New(httpServer, ":8080")

	// run app
	app.Run()

	// mock uuid library
	uuid.SetRand(nil)

	cases := []struct {
		name         string
		method       string
		endpoint     string
		body         string
		wantStatus   int
		wantResponse string
	}{
		{
			name:         "Create Booking - Valid Request",
			method:       http.MethodPost,
			endpoint:     "/bookings",
			body:         `{"guest_id":"g_123","pickup":{"query":"Barcelona El Prat T1"},"dropoff":{"query":"Passeig de Gràcia 10, Barcelona"},"pickup_at":"2026-02-01T10:00:00Z00:00"}`,
			wantStatus:   http.StatusCreated,
			wantResponse: `{"booking_id": "some-uuid", "estimation": 99.00}`,
		},
	}

	// run cases
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// send request

			req := httptest.NewRequest(tc.method, tc.endpoint, strings.NewReader(tc.body))

			recorder := httptest.NewRecorder()

			httpServer.ServeHTTP(recorder, req)

			// check response

			res := recorder.Result()
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			assert.NoError(t, err)

			assert.Equal(t, tc.wantStatus, res.StatusCode)
			assert.Equal(t, tc.wantResponse, string(body))

		})
	}

}
