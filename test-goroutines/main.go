package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world!")

	// useContext()
	useChannel()

	fmt.Println("Goodbye!")
}

func useChannel() {
	fmt.Println("use channel")
	ch := make(chan int, 1)

	go workerUseChannel(ch)

	time.Sleep(1 * time.Second)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 3
	ch <- 4
	close(ch)
	time.Sleep(2 * time.Second)
}

func useContext() {
	fmt.Println("use context")
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	go worker(ctx)

	time.Sleep(1 * time.Second)
	cancel()

	time.Sleep(2 * time.Second)
}

func worker(ctx context.Context) {
	fmt.Println("Worker started")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker stopped")
			return
		default:
			fmt.Println("Worker working...")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Worker finish")
		}
	}

}

func workerUseChannel(ch chan int) {
	for {
		select {
		case <-ch:
			fmt.Println("worker stopped")
			fmt.Println(ch, <-ch)
			fmt.Println("worker stopped")
			return

		default:
			fmt.Println("worker working...")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("worker finish")
		}
	}
}
