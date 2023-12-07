package main

import (
	"fmt"
	"math/rand"
	"time"
)

func basic(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	// after run this line, the main goroutine is finished.
	// main goroutine is a caller. It doesn't wait for func basic finished
	// Thus, we don't see anything
	go basic("basic!") // spawn a goroutine. (1)

	// To solve it, we can make the main go routine run forever by `for {}` statement.

	// for {
	// }

	// A little more interesting is the main goroutine exit. the program also exited
	// This code hang
	fmt.Println("This is the beginning of basic implementation")
	time.Sleep(2 * time.Second)
	fmt.Println("You're basic. I'm leaving")

	// However, the main goroutine and basic goroutine does not communicate each other.
	// Thus, the above code is cheated because the basic goroutine prints to stdout by its own function.
	// the line `basic! 1` that we see on terminal is the output from basic goroutine.

	// real conversation requires a communication
}
