package main

import (
    "fmt"
    "time"
)

// Example: a buffered channel lets you send up to N messages without an immediate receiver.
func main() {
    ch := make(chan struct{}, 3) // buffer for 3 messages

    // We can send 3 signals even if the worker isn't ready yet.
    ch <- struct{}{}
    ch <- struct{}{}
    ch <- struct{}{}
    fmt.Println("main: sent 3 buffered signals")

    go func() {
        // Consume the 3 signals
        for i := 0; i < 3; i++ {
            <-ch
            fmt.Printf("worker: received signal %d\n", i+1)
            time.Sleep(200 * time.Millisecond)
        }
        fmt.Println("worker: done")
    }()

    time.Sleep(1 * time.Second)
    fmt.Println("main: done")
}
