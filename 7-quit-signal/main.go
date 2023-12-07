package main

import (
	"fmt"
	"math/rand"
	"time"
)

// the basic function return a channel to communicate with it.
func basic(msg string, quit chan string) <-chan string { // <-chan string means receives-only channel of string.
	c := make(chan string)
	go func() { // we launch goroutine inside a function.
		for i := 0; i < 5; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				// do nothing
			case <-quit:
				fmt.Println("clean up")
				quit <- "See you!"
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

	}()
	return c // return a channel to caller
}

func main() {
	quit := make(chan string)
	c := basic("Joe", quit)
	for i := 3; i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye"
	fmt.Println("Joe say:", <-quit)
}
