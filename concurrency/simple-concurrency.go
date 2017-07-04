package main

import "fmt"

// Examine the basic principle of concurrency in Go using unbuffered channels and goroutines

func main() {
	// intChannel := make(chan int) // make an unbuffered int channel
	// stringChannel := make(chan string) // make an unbuffered string channel
	// interfaceChannel := make(chan interface{}) // make an unbuffered interface channel

	intSlice := []int{3, 15, 2, 8, 92, 91, 64} // create an int slice with a length of 7
	intChannel := make(chan int)
	go sum(intSlice[:len(intSlice)/2], intChannel) // spawn a goroutine to sum the first-half of the numbers in the intSlice ([:3]); sends the summed result to intChannel
	go sum(intSlice[len(intSlice)/2:], intChannel) // spawn a goroutine to sum the second-half of the numbers in the intSlice ([3:]); sends the summed result to intChannel
	x, y := <-intChannel, <-intChannel             // "receive" the sum value from intChannel and assign to a variable; the order of assignment here is not guaranteed...

	fmt.Println("The first half of intSlice is:", x, "\nThe second half of intSlice is:", y, "\nThe total sum of intSlice is:", x+y)
}

// sum computes the sum of an int slice.
// Sends the summed value to the given channel.
func sum(n []int, c chan int) {
	total := 0
	for _, v := range n {
		total += v
	}
	c <- total // "send" the value to the channel
}
