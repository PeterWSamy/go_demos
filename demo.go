package main

import (
	"fmt"
	"time"
)

func unbufferedChannelDemo() {
	fmt.Println("Unbuffered Channel Demo")
	ch := make(chan int)

	go func() {
		fmt.Println("Sending 1 to the unbuffered channel")
		ch <- 1
		fmt.Println("Sent 1 to the unbuffered channel")

	}()

	fmt.Println("Receiving from the unbuffered channel")
	value := <-ch
	fmt.Println("Received", value, "from the unbuffered channel")
}

func bufferedChannelDemo() {
	fmt.Println("Buffered Channel Demo")
	ch := make(chan int, 2)

	go func() {
		fmt.Println("Sending 1 to the buffered channel")
		ch <- 1 // This will not block
		fmt.Println("Sent 1 to the buffered channel")

		fmt.Println("Sending 2 to the buffered channel")
		ch <- 2 // This will not block
		fmt.Println("Sent 2 to the buffered channel")

		fmt.Println("Sending 3 to the buffered channel   " + time.Now().String())
		ch <- 3 // This will block
		fmt.Println("Sent 3 to the buffered channel")
	}()

	time.Sleep(time.Second) // Simulate some work

	fmt.Println("Receiving from the buffered channel")
	value1 := <-ch
	fmt.Println("Received", value1, "from the buffered channel")

	fmt.Println("Receiving from the buffered channel")
	value2 := <-ch
	fmt.Println("Received", value2, "from the buffered channel   ", time.Now().String())
}

func main() {
	unbufferedChannelDemo()
	fmt.Println()
	bufferedChannelDemo()
}
