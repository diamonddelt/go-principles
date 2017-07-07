package main

import "fmt"

func doubler(c, stop chan int) {
	x := 1
	for { // infinitely monitor for and execute the following
		select { // select will act like a switch/case statement, but will choose the first-available channel sending/receiving
		case c <- x: // choose this case when the c channel is not locked, and fill it with the current value of x
			x = x * 2
		case <-stop:
			fmt.Println("Exiting...")
			return // break out of the infinite loop
		default:
			// fmt.Println("Uncomment me to see how many cycles of the infinite loop in doubler() occured before either of the cases were available.")
		}
	}
}

func main() {
	c := make(chan int)    // make an unbuffered channel 'c'
	stop := make(chan int) // make an unbuffered channel 'stop' to signal stopping the doubler function
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // 'pop' or receive a piece of data from the channel 'c' 10 times
		}
		stop <- 0 // signal the 'stop' channel case in doubler()
	}()
	doubler(c, stop)
}
