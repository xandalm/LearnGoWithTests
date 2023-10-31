package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buf := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buf, spySleeper)

	got := buf.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enought calls to sleeper, want 3 got %d", spySleeper.Calls)
	}
}
