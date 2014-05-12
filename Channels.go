// Channels
package main

import (
	"fmt"
	"time"
)

func channels_() {
	messages := make(chan string)
	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}

func channels_buffering() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func worker(done chan bool) {
	fmt.Print("working ...")
	time.Sleep(time.Second)
	fmt.Println("done.")

	done <- true
}

func channel_synchronization() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func channel_directions() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message.")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

/*
Basic sends and receives on channels are blocking.
However, we can use select with a default clause to
implement non-blocking sends, receives, and
even non-blocking multi-way selects.
*/
func channel_non_blocking() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("send message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

func channel_close() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			//In this special 2-value form of receive,
			//the more value will be false if jobs
			//has been closed and all values
			//in the channel have already been received.
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
		time.Sleep(time.Second * 1)
	}
	/*
		chan关闭以后，还是会无阻塞的收到空数值(0),
		所以在<-时需要通过第二个参数判断chan是否关闭.
		语法为: value, ok := <-chan,if !ok,chan closed
	*/
	close(jobs)
	fmt.Println("send all jobs")
	<-done
}

func channel_range() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	/*
		This range iterates over each element
		as it is received from queue.
		Because we closed the channel above,
		the iteration terminates after
		receiving the 2 elements.
		If we did not close it we would block
		on a 3rd receive in the loop.

		This example also showed that
		it is possible to close a non-empty
		channel but still have the remaining
		values be received.
	*/
	for elem := range queue {
		fmt.Println(elem)
	}
}
