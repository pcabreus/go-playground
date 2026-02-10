package main

import (
	"fmt"
	"net/http"

	"github.com/pcabreus/challenges/bookingaride/internal"
)

func main() {

	// Load config

	mux := http.NewServeMux()

	app := internal.New(mux, ":8080")
	// app - Rouiting system, handlers, services, statorage

	// run application
	if err := app.Run(); err != nil {
		fmt.Println("Application error:", err)
	}
}
