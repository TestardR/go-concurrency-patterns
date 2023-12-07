package main

import (
	"fmt"
	"math/rand"
	"time"
)

// now, the basic function additional parameter
// `c chan string` is a channel
func basic(c chan string, msg string) {
	for i := 0; i < 5; i++ {
		// send the value to channel
		// it also waits for receiver to be ready
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	// init our channel
	c := make(chan string)
	go basic(c, "basic!")

	for i := 0; i < 5; i++ {
		// <-c read the value from basic function
		// <-c waits for a value to be sent
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're basic. I'm leaving")
}
