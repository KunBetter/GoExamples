// RateLimiting
package main

import (
	"fmt"
	"time"
)

func RateLimiting_() {
	p := fmt.Println
	request := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		request <- i
	}
	close(request)

	/*
		This limiter channel will receive a value
		every 200 milliseconds. This is the regulator
		in our rate limiting scheme.
	*/
	limiter := time.Tick(time.Millisecond * 200)
	/*
		By blocking on a receive from the limiter channel
		before serving each request,
		we limit ourselves to 1 request every 200 milliseconds.
	*/
	for req := range request {
		<-limiter
		p("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		p("request", req, time.Now())
	}
}
