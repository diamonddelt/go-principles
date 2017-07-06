package main

import "fmt"

func main() {
	c := make(chan int, 5) // create an int channel with a buffered capacity of 3 for use with factorial()
	go counter(cap(c), c)  // spawn a goroutine to run factorial(); cap() specifies the capacity of the channel

	count := 1
	for i := range c { // this line is important - without the explicit close(c) in counter(), this range over a channel would deadlock
		fmt.Println("The current factorial value at iteration", count, "is:", i)
		count++
	}
}

func counter(n int, c chan int) {
	count := 1
	for i := 0; i < n; i++ {
		c <- count // send the current count to the channel
		count++
	}
	close(c) // explicitly close the channel; this is needed for goroutine access-safety in the caller
	// "close should always be executed by the sender, never the receiver" <- https://golang.org/pkg/builtin/#close
}
