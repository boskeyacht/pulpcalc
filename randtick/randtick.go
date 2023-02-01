package randtick

import (
	"math/rand"
	"time"
)

type RandTick struct {
	// The ticker channel
	C chan time.Time

	// A channel to stop on
	stop chan chan struct{}

	// The interval's lower bound
	Min int64

	// The interval's upper bound
	Max int64
}

// Creates a new `RandTick`er with an interval between `min` and `max`
func NewRandTick(min, max time.Duration) *RandTick {
	return &RandTick{
		C:    make(chan time.Time),
		stop: make(chan chan struct{}),
		Min:  min.Nanoseconds(),
		Max:  max.Nanoseconds(),
	}
}

// Creates a new `RandTick`er with an interval between 0 and `max`
func NewRandTickN(max time.Duration) *RandTick {
	return &RandTick{
		C:    make(chan time.Time),
		stop: make(chan chan struct{}),
		Min:  time.Duration(1).Nanoseconds(),
		Max:  max.Nanoseconds(),
	}
}

func (r *RandTick) Stop() {
	c := make(chan struct{})
	r.stop <- c
	<-c
}

func (r *RandTick) Start() {
	defer close(r.C)
	t := time.NewTimer(r.tick())

	for {
		select {
		case <-t.C:
			select {
			case r.C <- time.Now():
				t.Stop()
				t = time.NewTimer(r.tick())
			}
		case c := <-r.stop:
			t.Stop()
			c <- struct{}{}
			return
		}
	}
}

func (r *RandTick) tick() time.Duration {
	i := rand.Int63n(int64(r.Max-r.Min)) + int64(r.Min)

	return time.Duration(i) * time.Nanosecond
}
