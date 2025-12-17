package main

import (
    "fmt"
    "sync"
    "time"
)

// Example: multiple receivers and a single send
// Demonstrates that a single send is delivered to only one receiver.
func main() {
    ch := make(chan struct{})
    var wg sync.WaitGroup
    wg.Add(2)

    worker := func(id int) {
        defer wg.Done()
        <-ch // espera una señal
        fmt.Printf("worker %d: recibió señal y termina\n", id)
    }

    go worker(1)
    go worker(2)

    time.Sleep(500 * time.Millisecond)
    fmt.Println("main: sending A SINGLE signal")
    ch <- struct{}{} // will be consumed by ONE worker

    // If we want to wake up or ensure all workers finish, we can close the channel
    // close(ch) // uncomment to wake everyone up via zero-value receives

    // Give some time and then close to avoid a permanent block if we don't send again
    time.Sleep(200 * time.Millisecond)
    select {
    case ch <- struct{}{}:
        // if there is a buffer or an extra receiver, this second send will reach another worker
    default:
        // do nothing
    }

    // If none of the workers picked up the second send, close so they do not remain blocked
    close(ch)

    wg.Wait()
    fmt.Println("main: done")
}
