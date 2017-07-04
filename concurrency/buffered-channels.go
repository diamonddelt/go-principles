package main

import "fmt"

// Examine the basic behavior of a buffered channel with a given capacity, as opposed to an unbuffered channel

func main() {
	bufferedChannel := make(chan int, 5)
	fillBufferedChannel(5, bufferedChannel)
	// fillBufferedChannel(6, bufferedChannel) // NOTE: this will cause a runtime panic in the main goroutine if you exceed the bufferedChannel's capacity (5)

	for i := 0; i < 5; i++ {
		fmt.Println("The value from the bufferedChannel is:", <-bufferedChannel) // receives data from the channel
	}

}

// fillBufferedChannel sends integers iterated through a for-loop to the given channel
func fillBufferedChannel(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i // send the increment of the for loop to the channel
	}
}
