package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord     = "Go!"
	countDownTime = 3
)

func CountDown(w io.Writer, sleeper Sleeper) {
	for i := countDownTime; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := ConfigurableSleeper{4 * time.Second, time.Sleep}
	CountDown(os.Stdout, sleeper)
}
