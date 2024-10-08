package channels

import "fmt"

func BufferedChannels() {
	dataChan := make(chan int, 1)
	var n int

	dataChan <- 1

	n = <-dataChan

	fmt.Printf("buffered n = %d\n", n)

}
