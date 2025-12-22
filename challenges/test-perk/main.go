package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type response struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(response{
		Message: "pong",
		Time:    time.Now().Format(time.RFC3339),
	})
}

func main() {
	// Load .env
	_ := godotenv.Load()

	token := getEnv("FOURSQUARE_API_KEY", "")
	add := getEnv("ADDR", ":8080")

	
	client := NewFourSquare("https://places-api.foursquare.com", token)

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingHandler)
	mux.HandleFunc("/places/search", func(w http.ResponseWriter, r *http.Request) { // query -> lat and long sort near response [{name, lat, long, distance}]
		lat := r.URL.Query().Get("lat")
		long := r.URL.Query().Get("long")
		// sort is near
		places, err := client.pointOfInterest(lat, long)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp, err := json.Marshal(places)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	})

	srv := &http.Server{
		Addr:              add,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("listening on %s\n", add)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}

}

type FourSquare struct {
	client *http.Client
	host   string
	apiToken string
}

type PlaceSearchResponse struct {
	Results []Places `json:"results"`
}

type Places struct {
	Name     string  `json:"name"`
	Lat      float64 `json:"latitude"`
	Long     float64 `json:"longitude"`
	Distance float64 `json:"-"`
}

func NewFourSquare(host,token string) FourSquare {
	client := &http.Client{
		Timeout: 5 * time.Second, // timeout total (incluye DNS + connect + body)
	}

	return FourSquare{
		client: client,
		host:   host,
		token: token,
	}
}

func (f FourSquare) pointOfInterest(lat, long string) ([]Places, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	path := fmt.Sprintf("%s/places/search?ll=%s,%s", f.host, lat, long)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("X-Places-Api-Version", "2025-06-17")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+)

	resp, err := f.client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err // TODO:
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err // TODO Wrap this
	}

	var placeSearchResponse PlaceSearchResponse
	err = json.Unmarshal(body, &placeSearchResponse)
	if err != nil {
		return nil, err // TODO
	}

	pointLat, _ := strconv.ParseFloat(lat, 32)
	pointLong, _ := strconv.ParseFloat(long, 32)
	for i, place := range placeSearchResponse.Results {
		distance := (place.Lat-pointLat)*(place.Lat-pointLat) + (place.Long-pointLong)*(place.Long-pointLong)

		placeSearchResponse.Results[i].Distance = distance
	}

	sort.Slice(placeSearchResponse.Results, func(i, j int) bool {
		return placeSearchResponse.Results[i].Distance > placeSearchResponse.Results[j].Distance
	})

	return placeSearchResponse.Results, err
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
