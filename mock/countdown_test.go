package main

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {

	t.Run("print 3 to go!", func(t *testing.T) {
		buf := &bytes.Buffer{}
		// spySleeper := &SpySleeper{}
		// Countdown(buf, spySleeper)
		Countdown(buf, &SpyCountdownOperations{})

		got := buf.String()
		want := "3\n2\n1\nGo!\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		/*
			if spySleeper.Calls != 3 {
				t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
			}
		*/
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}

}
