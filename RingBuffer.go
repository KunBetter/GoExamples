// RingBuffer
package main

import (
	"fmt"
)

/*
	Connect two buffered channels through one goroutine
	that forwards messages from the incoming channel
	to the outgoing channel.
	Whenever a new message can not be placed
	on the outgoing channel,
	take one message out of the outgoing channel
	(that is the oldest message in the buffer),
	drop it, and place the new message
	in the newly freed up outgoing channel.

	Plugging in this “channel struct” will never block
	and will simply behave like a ring buffer.
	That is, slower consumers might loose (their oldest)
	messages, but will never be able to
	block the main message processing loop.
*/

type RingBuffer struct {
	inputChannel  <-chan int
	outputChannel chan int
}

func NewRingBuffer(inputChannel <-chan int, outputChannel chan int) *RingBuffer {
	return &RingBuffer{inputChannel, outputChannel}
}

func (r *RingBuffer) Run() {
	for v := range r.inputChannel {
		select {
		case r.outputChannel <- v:
		default:
			<-r.outputChannel
			r.outputChannel <- v
		}
	}
	close(r.outputChannel)
}

func main() {
	in := make(chan int)
	out := make(chan int, 6)
	rb := NewRingBuffer(in, out)
	go rb.Run()

	for i := 0; i < 10; i++ {
		in <- i
	}

	close(in)

	for res := range out {
		fmt.Println(res)
	}
}
