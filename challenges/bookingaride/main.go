package main

import "fmt"

func main() {

	// Load config

	app := internal.App{}
	// app - Rouiting system, handlers, services, statorage

	// run application
	if err := app.Run(); err != nil {
		fmt.Println("Application error:", err)
	}
}
