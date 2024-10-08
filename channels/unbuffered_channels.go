package channels

import "fmt"

// why error :=
/*
	Issue: Since the channel is unbuffered, it requires a receiver ready to consume the value immediately.
	However, the code is trying to send without having a receiver in place, so it blocks here.
*/

func UnBufferedChannelsError() {
	dataChan := make(chan int)
	var n int

	go func() {
		n = <-dataChan

	}()

	dataChan <- 1

	fmt.Printf("n = %d\n", n)

}

// How to Solve :
/*
  When you try to send a value on an unbuffered channel, the send operation will block (pause execution) until a receiver is ready to receive the value.
	•	Similarly, when you try to receive from an unbuffered channel, the receive operation will block until a sender sends a value.

  In other words, both the send and receive must happen at the same time for the communication to succeed.
  Result : we ready receiver another go routine
*/

func UnBufferedChannelsFix() {
	dataChan := make(chan int)
	var n int

	go func() {
		n = <-dataChan

	}()

	dataChan <- 1

	fmt.Printf("n = %d\n", n)

}
