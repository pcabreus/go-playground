package main

import (
    "fmt"
    "time"
)

// Example: close a channel and use range to process values until the channel is closed.
func main() {
    ch := make(chan struct{})

    go func() {
        for range ch {
            fmt.Println("worker: processing signal")
            time.Sleep(150 * time.Millisecond)
        }
        fmt.Println("worker: channel closed, exiting")
    }()

    // Send 3 signals and then close the channel to indicate there will be no more.
    ch <- struct{}{}
    ch <- struct{}{}
    ch <- struct{}{}

    close(ch)

    time.Sleep(1 * time.Second)
    fmt.Println("main: done")
}
