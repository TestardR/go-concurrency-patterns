package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func fanIn(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for i := range inputs {
		input := inputs[i]
		go func() {
			for {
				c <- <-input
			}
		}()
	}
	return c
}

// the basic function return a channel to communicate with it.
func basic(msg string) <-chan Message { // <-chan Message means receives-only channel of Message.
	c := make(chan Message)
	waitForIt := make(chan bool) // share between all messages
	go func() {                  // we launch goroutine inside a function.
		for i := 0; i < 5; i++ {
			c <- Message{
				str:  fmt.Sprintf("%s %d", msg, i),
				wait: waitForIt,
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

			// every time the goroutine send message.
			// This code waits until the value to be received.
			<-waitForIt
		}

	}()
	return c // return a channel to caller.
}

func main() {
	// merge 2 channels into 1 channel
	c := fanIn(basic("john"), basic("doe"))

	for i := 0; i < 5; i++ {
		msg1 := <-c // wait to receive message
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)

		// each go routine have to wait
		msg1.wait <- true // main goroutine allows the basic goroutine to send next value to message channel.
		msg2.wait <- true
	}
	fmt.Println("You're both basic. I'm leaving")
}

// main: goroutine                                          basic: goroutine
//    |                                                           |
//    |                                                           |
// wait for receiving msg from channel c                    c <- Message{} // send message
//   <-c                                                          |
//    |                                                           |
//    |                                                     <-waitForIt // wait for wake up signal
// send value to channel                                          |
// hey, basic. You can send next value to me                     |
//   wait <-true                                                  |
///                            REPEAT THE PROCESS
