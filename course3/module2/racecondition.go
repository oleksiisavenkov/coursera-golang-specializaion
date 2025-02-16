/* Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur. */

/*
	The provided Go code demonstrates a simple program that highlights a potential race condition. The program consists of a global variable i, a function sendToChannelAndIncrement, and the main function.

	Function: sendToChannelAndIncrement
	The sendToChannelAndIncrement function takes a channel c as an argument. It sends the current value of the global variable i to the channel and then increments i by one. This function is intended to be run concurrently by multiple goroutines.

	Function: main
	In the main function, a channel c is created to facilitate communication between goroutines. A loop is used to start 10 goroutines, each executing the sendToChannelAndIncrement function. This means that 10 concurrent operations will attempt to send the value of i to the channel and then increment it.

	After starting the goroutines, another loop is used to receive values from the channel. The select statement is used to either:

	Receive a value from the channel and print it using fmt.Println, or
	Timeout after 1 second using time.After. If the timeout occurs, the program prints "Timeout" and exits.
	Race Condition Explanation
	The race condition arises because multiple goroutines access and modify the global variable i concurrently without synchronization. This can lead to unpredictable behavior, as the value of i may not be incremented correctly.

	To avoid this issue, synchronization mechanisms such as mutexes (sync.Mutex) should be used to ensure that only one goroutine can access and modify i at a time.
*/

package main

import (
	"fmt"
	"time"
)

var i int = 0

func sendToChannelAndIncrement(c chan int) {
	c <- i
	i++
}

func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go sendToChannelAndIncrement(c)
	}
	for i := 0; i < 10; i++ {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout")
			return
		}
	}
}
