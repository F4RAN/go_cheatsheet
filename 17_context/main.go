package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1. context.Background()
	// Root context, usually created in main.
	// Application: root context of the app.
	// Used in: main(), app startup, tests, first/base context.
	ctx1 := context.Background()
	fmt.Println("1. background context:", ctx1)

	// 2. context.WithCancel()
	// Manual cancel signal.
	// Application: manually stop a worker.
	// Used in: graceful shutdown, stopping goroutines, canceling background jobs.
	ctx2, cancel2 := context.WithCancel(context.Background())

	go func() {
		<-ctx2.Done()
		fmt.Println("2. canceled manually")

	}()
	// defer runs later, right before the current function returns.
	cancel2()
	time.Sleep(100 * time.Millisecond)

	// 3. context.WithTimeout()
	// Auto cancel after a duration.
	// Application: cancel work if it takes too long.
	// Used in: API calls, database queries, HTTP requests, network operations.
	ctx3, cancel3 := context.WithTimeout(context.Background(), 1*time.Second)
	// Multiple defer calls run in reverse order.
	defer cancel3()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("3. work finished")
	case <-ctx3.Done():
		fmt.Println("3. timeout:", ctx3.Err())
	}

	// 4. context.WithDeadline()
	// Auto cancel at exact time.
	// Application: cancel at a specific time.
	// Used in: scheduled jobs, request expiry, operations that must finish before exact time.
	deadline := time.Now().Add(1 * time.Second)
	ctx4, cancel4 := context.WithDeadline(context.Background(), deadline)
	defer cancel4()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("4. work finished")
	case <-ctx4.Done():
		fmt.Println("4. deadline reached:", ctx4.Err())
	}

	// 5. context.WithValue()
	// Pass request-scoped data.
	// Application: pass request metadata.
	// Used in: request ID, trace ID, user ID, auth/session metadata, logging context.
	ctx5 := context.WithValue(context.Background(), "request_id", "abc-123")

	requestID := ctx5.Value("request_id")
	fmt.Println("5. request id:", requestID)

}
