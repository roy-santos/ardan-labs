package main

import (
	"fmt"
	"time"
)

func main() {

	// this is how you start a goroutine with "go" keyword, once it starts you can't access it anymore
	// defer does not wait for the goroutine either
	go fmt.Println("goroutine")

	fmt.Println("main")

	for i := 0; i < 3; i++ {

		/* BUG: All goroutines use the "i" for the for loop
		go func() {
			fmt.Println(i) // this uses "i" from line 16, prints out 3 because i == 3 when it starts goroutine
		}()
		*/

		/* Fix 1: use a parameter
		go func(n int) {
			fmt.Println(n)
		}(i)
		*/

		// Fix 2: Use a loop body variable
		i := i // "i" shadows "i" from the for loop, see that "i" escapes to the heap if you run "go build -gcflags=-m"
		go func() {
			fmt.Println(i) // i from line 31
		}()

		/* Can also do this to avoid the whole shadowing thing, not sure why this works tho..
		func() {
			go fmt.Println(i)
		}()
		*/

	}

	time.Sleep(10 * time.Millisecond)

	shadowExample()

	/*
		Channels are not like queues where if nobody receives sent message, it will move on. In a channel, a sent message
		must be received before the code can continue processing
	*/
	ch := make(chan string) // this is making a channel
	go func() {
		ch <- "hi" // send
	}()
	msg := <-ch // receive
	fmt.Println(msg)

	// this is the producer that generates messages
	// messages are consumed in the main fxn
	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("message #%d", i+1)
			ch <- msg // sends msg to the channel
		}
		close(ch) // this will allow range fxn on the channel, range knows where to stop
	}()

	// doing a range over a channel, go will not know when there are no messages left, this causes error
	for msg := range ch {
		fmt.Println("got:", msg)
	}

	/* above for/range basically does this..
	for {
		msg, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("got:", msg)
	}
	*/

	msg = <-ch // ch is already closed
	fmt.Printf("closed: %#v\n", msg)

	msg, ok := <-ch // using the ok will tell you whether the channel is closed (you'll receive 0 value) or if it actually sent that 0
	fmt.Printf("closed: %#v (ok=%v)\n", msg, ok)

	// ch <- "hi" // ch is closed, this will panic

	values := []int{15, 8, 42, 16, 4, 23}
	fmt.Println(sleepSort(values))
}

/*
For every value "n" in values, spin a goroutine that will
- sleep "n" milliseconds
- send "n" over a channel

I the funciton body, collect values from the channel to a slice and return it.
*/
func sleepSort(values []int) []int {
	ch := make(chan int)

	for _, n := range values {
		n := n
		go func() {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n
		}()
	}

	var out []int
	for range values {
		n := <-ch
		out = append(out, n)
	}

	return out
}

/* Channel semantics
- send & receive will block until opposite operation (*, there is an exception)
- receive from a closed channel will return the zero value without blocking
- send to a closed channel will panic
- only owner of the channel can close it and send to it
- closing a closed channel will panic
- don't have to close a channel, only if you want to notify that nothing more is coming
*/

// see also: https://www.353solutions.com/channel-semantics

func shadowExample() {
	n := 7
	{
		n := 2 // n == 2 from here to }
		fmt.Println("inner", n)
	}
	fmt.Println("outer", n)
}
