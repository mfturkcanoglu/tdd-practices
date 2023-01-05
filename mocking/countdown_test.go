package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountDown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpyCountdownOperations{}
		CountDown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("expected %q but got %q", want, got)
		}

	})

	t.Run("sleep before every sleep", func(t *testing.T) {
		spySleeper := &SpyCountdownOperations{}
		CountDown(spySleeper, spySleeper)

		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(spySleeper.Calls, want) {
			t.Errorf("expected %q but got %q", want, spySleeper.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
