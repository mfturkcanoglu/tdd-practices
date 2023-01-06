package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondsTimeout = 10 * time.Second

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		winner = a
	case <-ping(b):
		winner = b
	case <-time.After(timeout):
		err = fmt.Errorf("timed out waiting for %q and %q", a, b)
	}
	return
}

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondsTimeout)
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func RacerWithoutConcurrency(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		winner = a
	} else {
		winner = b
	}
	return
}

func measureResponseTime(url string) (duration time.Duration) {
	startTime := time.Now()
	http.Get(url)
	duration = time.Since(startTime)
	return
}
