package main

import (
	"fmt"
	"math/rand"
	"time"
)

// basic is a function that returns a channel to communicate with it.
// <-chan string means receives-only channel of string.
func basic(msg string) <-chan string {
	c := make(chan string)
	// we launch goroutine inside a function
	// that sends the data to channel
	go func() {
		// The for loop simulate the infinite sender.
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

		// The sender should close the channel
		close(c)

	}()
	return c // return a channel to caller.
}

func main() {

	john := basic("john")
	doe := basic("doe")

	// This loop yields 2 channels in sequence
	for i := 0; i < 5; i++ {
		fmt.Println(<-john)
		fmt.Println(<-doe)
	}

	// or we can simply use the for range
	// for msg := range john {
	// 	fmt.Println(msg)
	// }
	fmt.Println("You're both basic. I'm leaving")

}
