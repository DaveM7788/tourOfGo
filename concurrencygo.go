package main

import (
	"fmt"
	"time"
)

func main() {
	simpleGoroutine()
	channelsChan()
	bufferedChan()
	rangeAndClose()
	selectSel()
	defaultSelect()
}

func simpleGoroutine() {
	// A goroutine is a lightweight thread managed by the Go runtime. uses go keyword
	go say("world")
	say("hello")
	// go f(x, y, z)
	// The evaluation of f, x, y, and z happens in the current goroutine and the execution
	// of f happens in the new goroutine
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// Channels are a typed conduit through which you can send and receive values with the channel operator, <-
func channelsChan() {
	s := []int{7, 2, 8, -9, 4, 0}

	// Like maps and slices, channels must be created before use. Note the chan keyword
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

/*
By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without
explicit locks or condition variables.

The example code above sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines
have completed their computation, it calculates the final result
*/

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func bufferedChan() {
	// Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel
	// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3  // would cause a deadlock. overfilling buffer
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*
A sender can close a channel to indicate that no more values will be sent.
Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression
v, ok := <-ch
*/
func rangeAndClose() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// note only senders should close channels. not receivers
	close(c)
	// you do not have to close channels. useful for termination purposes
}

/*
The select statement lets a goroutine wait on multiple communication operations.

A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
*/
func selectSel() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciSel(c, quit)
}

func fibonacciSel(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// The default case in a select is run if no other case is ready.
func defaultSelect() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
