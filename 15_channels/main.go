package main

import "fmt"

func main() {
	// 1. Unbuffered channel
	ch1 := make(chan int)
	go func() { ch1 <- 1 }()
	fmt.Println(<-ch1)

	// 2. Buffered channel
	ch2 := make(chan int, 2)
	ch2 <- 2
	fmt.Println(<-ch2)
	ch2 <- 3
	fmt.Println(<-ch2)

	// 3. Send-only channel
	ch3 := make(chan int)
	go func(ch chan<- int) { ch <- 4 }(ch3) // Expanded version is at the end (sendOnly func)
	fmt.Println(<-ch3)

	// 4. Receive-only channel
	ch4 := make(chan int, 1)
	ch4 <- 5
	func(ch <-chan int) { // Expanded version is at the end (receiveOnly func)
		fmt.Println(<-ch)
	}(ch4)

	// 5. Nil channel
	var ch5 chan int
	fmt.Println(ch5 == nil)

	// 6. Closed channel
	ch6 := make(chan int, 1)
	ch6 <- 6
	close(ch6)

	value, ok := <-ch6
	fmt.Println("closed channel value:", value, "ok:", ok)

	value, ok = <-ch6
	fmt.Println("closed channel after empty:", value, "ok:", ok)

	// 7. Channel of channels
	ch7 := make(chan chan int)
	go func() {
		inner := make(chan int, 1)
		inner <- 7
		ch7 <- inner
	}()
	fmt.Println(<-(<-ch7))

	// 8. bool channel: signal + true/false value
	ch8 := make(chan bool)
	go func() {
		ch8 <- true
	}()
	fmt.Println("bool channel:", <-ch8)

	// 9. number channel: signal + numeric value/code
	ch9 := make(chan int)
	go func() {
		ch9 <- 200
	}()
	fmt.Println("number channel:", <-ch9)

	// 10. struct{} channel: pure signal only,
	// no data, good for signaling, memory usage 0
	ch10 := make(chan struct{})
	go func() {
		close(ch10)
	}()
	<-ch10
	fmt.Println("struct{} channel: done")

}

// EXAMPLE OF Send Only and Receive Only go routines
// func sendOnly(ch chan<- int) {
// 	ch <- 100
// }

// func receiveOnly(ch <-chan int) {
// 	value := <-ch
// 	fmt.Println("receive-only channel:", value)
// }
