package main

import (
	"fmt"
	"math/rand"
	"time"
)

// the basic function return a channel to communicate with it.
func basic(msg string) <-chan string { // <-chan string means receives-only channel of string.
	c := make(chan string)
	go func() { // we launch goroutine inside a function.
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		}

	}()
	return c // return a channel to caller.
}

func main() {
	c := basic("Joe")

	// timeout for the whole conversation
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much.")
			return
		}
	}
}
