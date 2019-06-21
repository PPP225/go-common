package common

import (
	"fmt"
	"log"
	"time"
)

// TimeRemainingTimer used to track progress
type TimeRemainingTimer struct {
	start     int64
	end       int64
	startTime int64
}

// CreateTimeRemainingTimer creates timer, that will interpolate remaining time based on some start & end indexes
func CreateTimeRemainingTimer(start, end int64) *TimeRemainingTimer {
	return &TimeRemainingTimer{
		start,
		end,
		time.Now().UnixNano(),
	}
}

// Start starts the timer
func (v *TimeRemainingTimer) Start() {
	v.startTime = time.Now().UnixNano()
}

// Restart restarts the timer with new start & end values
func (v *TimeRemainingTimer) Restart(start, end int64) {
	v.start = start
	v.end = end
	v.startTime = time.Now().UnixNano()
}

// Get returns %done & remainingTime for current index
func (v *TimeRemainingTimer) Get(current int64) (percentageDone float64, timeRemaining time.Duration) {
	now := time.Now().UnixNano()
	elapsed := now - v.startTime

	done := current - v.start
	total := v.end - v.start

	resultPerc := float64(done) / float64(total)
	remainingTime := int64(0)
	if resultPerc != 0 {
		remainingTime = int64(float64(elapsed)/resultPerc - float64(elapsed))
	}
	resultRemainingTime := time.Duration(remainingTime)

	return resultPerc, resultRemainingTime
}

// Getf returns formatted string of %done & remaining time
func (v *TimeRemainingTimer) Getf(current int64) string {
	p, r := v.Get(current)
	return fmt.Sprintf("Perc: %5.2f | Rem: %s", p*100, r.Round(time.Millisecond*10))
}

// Printr fmt.Printf's formatted %done & remaining time
func (v *TimeRemainingTimer) Printr(current int64) {
	p, r := v.Get(current)
	fmt.Printf("\rPerc: %5.2f | Rem: %s |     ", p*100, r.Round(time.Millisecond*10))
}

/////////////////////////////////////////////////////////////////////////

// TimeTrack logs elapsed time for task
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed.Round(time.Millisecond*10))
}
