package main

import "fmt"

func main() {
	c := make(chan int, 5)  // create an int channel with a buffered capacity of 3 for use with factorial()
	go factorial(cap(c), c) // spawn a goroutine to run factorial(); cap() specifies the capacity of the channel

	count := 1
	for i := range c {
		fmt.Println("The current factorial value at iteration", count, "is:", i)
		count++
	}
}

func factorial(n int, c chan int) {
	f := 1
	for i := 1; i < (n + 1); i++ {
		c <- f // send the current factorial value to the channel
		f *= i // get new factorial value
	}
	close(c) // explicitly close the channel; this is needed for goroutine access-safety in the caller
	// "close should always be executed by the sender, never the receiver" <- https://golang.org/pkg/builtin/#close
}
