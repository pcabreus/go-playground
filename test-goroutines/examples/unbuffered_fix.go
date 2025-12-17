package main

import (
    "fmt"
    "time"
)

// Example: stop signal with an unbuffered channel using close
// Avoids performing multiple receives that could create unexpected deadlocks.
func main() {
    ch := make(chan struct{})

    go func() {
        // The worker waits for the signal and then returns.
        <-ch
        fmt.Println("worker: received stop signal (close)")
    }()

    time.Sleep(1 * time.Second)
    // Close the channel: any remaining receives will return the zero value
    // immediately and the goroutine can proceed to exit.
    close(ch)

    // Give the worker a moment to print and finish.
    time.Sleep(300 * time.Millisecond)
    fmt.Println("main: done")
}
