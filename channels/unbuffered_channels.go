package channels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

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

	fmt.Printf("un buffered n = %d\n", n)

}

func UnBufferedChannelsForERR() {

	dataChan := make(chan int)

	go func() {
		for i := 0; i < 1000; i++ {
			dataChan <- i
		}
	}()

	for v := range dataChan {
		fmt.Printf("un buffered n = %d\n", v)
	}

}

/*
The Issue:

In an unbuffered channel, the sender (dataChan <- 1) will block until the receiver reads from the channel. In your code:

	1.	The goroutine is sending values into the channel.
	2.	The main goroutine is trying to range over the channel to receive values.
	3.	Deadlock happens because you never close the channel,
	so the for v := range dataChan keeps waiting for more values,
	but the sender goroutine eventually finishes after sending 1000 values and there’s no more data,
	leading to a situation where both are waiting indefinitely.
*/

/*

The deadlock occurs because the range loop over the channel does not stop unless the channel is closed.
Here’s the sequence of events leading to deadlock:

*/

func UnBufferedChannelsFor() {

	dataChan := make(chan int)

	go func() {
		for i := 0; i < 1000; i++ {
			dataChan <- i
		}
		close(dataChan)
	}()

	for v := range dataChan {
		fmt.Printf("un buffered n = %d\n", v)
	}

}

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func UnBufferedChannelsDoWork() {

	dataChan := make(chan int)

	go func() {
		for i := 0; i < 1000; i++ {
			// go func() {
			result := DoWork()
			dataChan <- result
			// }()
		}
		close(dataChan)
	}()

	for v := range dataChan {
		fmt.Printf("un buffered n = %d\n", v)
	}

}

func UnBufferedChannelsDoWork2() {

	dataChan := make(chan int)

	go func() {
		for i := 0; i < 1000; i++ {
			go func() {
				result := DoWork()
				dataChan <- result
			}()
		}
		close(dataChan)
	}()

	for v := range dataChan {
		fmt.Printf("un buffered n = %d\n", v)
	}

}

func UnBufferedChannelsDoWork3() {

	dataChan := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		}
		wg.Wait()
		close(dataChan)
	}()

	for v := range dataChan {
		fmt.Printf("un buffered n = %d\n", v)
	}

}
