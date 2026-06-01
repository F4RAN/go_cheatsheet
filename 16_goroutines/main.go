package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("1. hello from goroutine")

}

func main() {
	// 1. Run normal function in background
	go sayHello()

	time.Sleep(100 * time.Millisecond)

	// 2. Anonymous goroutine
	go func() {
		fmt.Println("2. anonymous goroutine")

	}()

	time.Sleep(100 * time.Millisecond)

	// 3. Goroutine with parameter
	go func(name string) { // Arguments
		fmt.Println("3. hello", name)

	}("Arian") // Parameter binding

	time.Sleep(100 * time.Millisecond)

	// 4. Goroutine for async task
	go func() {
		fmt.Println("4. sending email...")
		time.Sleep(1 * time.Second)
		fmt.Println("4. email sent")

	}()

	fmt.Println("4. main continues without waiting")

	time.Sleep(2 * time.Second)

	// 5. Multiple goroutines
	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Println("5. worker", id, "started")

		}(i)
	}

	fmt.Println("5. main is going ahead")

	time.Sleep(100 * time.Millisecond)

	// 6. Goroutine + channel to wait for result
	done := make(chan bool)

	go func() {
		fmt.Println("6. background job running")
		time.Sleep(1 * time.Second)
		done <- true
	}()

	<-done
	fmt.Println("6. background job finished")

}
