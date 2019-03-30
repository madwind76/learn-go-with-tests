package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Sleeper ...
type Sleeper interface {
	Sleep()
}

// SpySleeper ...
type SpySleeper struct {
	Calls int
}

// Sleep ...
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// DefaultSleeper ...
type DefaultSleeper struct{}

// Sleep ...
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// CountdownOperationsSpy ...
type CountdownOperationsSpy struct {
	Calls []string
}

const write = "write"
const sleep = "sleep"

// Sleep ...
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

// ConfigurableSleeper ...
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep ...
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// SpyTime ...
type SpyTime struct {
	durationSlept time.Duration
}

// Sleep ...
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// Countdown ...
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	// for i := countdownStart; i > 0; i-- {
	//	fmt.Fprintln(out, i)
	// }

	sleeper.Sleep()
	_, _ = fmt.Fprint(out, finalWord)

}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
