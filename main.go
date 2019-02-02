package main

import (
	"fmt"
	"syscall"
	"time"
	"unicode"

	"golang.org/x/crypto/ssh/terminal"
)

type stopwatch struct {
	elapsed     time.Duration
	lastElapsed time.Duration
	start       time.Time
	lastStart   time.Time
	isRunning   bool
}

func (s *stopwatch) Start() {
	if !s.isRunning {
		s.start = time.Now()
		s.lastStart = s.start
		s.isRunning = true
	}
}

func (s *stopwatch) Stop() {
	if s.isRunning {
		s.lastElapsed = time.Since(s.lastStart)
		s.elapsed += time.Since(s.start)
		s.isRunning = false
	}
}

func (s *stopwatch) Flip() {
	if s.isRunning {
		s.Stop()
	} else {
		s.Start()
	}
}

func (s *stopwatch) Elapsed() time.Duration {
	if s.isRunning {
		s.elapsed += time.Since(s.start)
		s.start = time.Now()
	}
	return s.elapsed
}

func newStopwatch(start bool) stopwatch {
	var sw stopwatch
	if start {
		sw.Start()
	}
	return sw
}

// Format Duration
func fd(d time.Duration) string {
	var ns string
	var period bool
	for _, r := range d.String() {
		if (period && unicode.IsDigit(r)) || r == '.' {
			period = true
			continue
		}
		ns += string(r)
	}
	return ns
}

func main() {
	working, distracted := newStopwatch(true), newStopwatch(false)

	fmt.Println("press enter to flip, ctrl+c to exit")
	fmt.Println("Type: Current Elapsed (Last Elapsed)")

	go func() {
		for {
			terminal.ReadPassword(int(syscall.Stdin))
			working.Flip()
			distracted.Flip()
		}
	}()

	for {
		fmt.Printf("\rWorking: %s (%s)\t\tDistracted %s (%s)\t\t", fd(working.Elapsed()), fd(working.lastElapsed), fd(distracted.Elapsed()), fd(distracted.lastElapsed))
		time.Sleep(time.Millisecond * 50)
	}
}
