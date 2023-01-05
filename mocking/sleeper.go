package main

import "time"

type Sleeper interface {
	Sleep()
}

type SpyCountdownOperations struct {
	Calls []string
}

func (spy *SpyCountdownOperations) Sleep() {
	spy.Calls = append(spy.Calls, sleep)
}

func (spy *SpyCountdownOperations) Write(bytes []byte) (n int, err error) {
	spy.Calls = append(spy.Calls, write)
	return
}

const (
	sleep = "sleep"
	write = "write"
)

type DefaultSleeper struct{}

func (def DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
