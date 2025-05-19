package timer

import (
	"time"
)

type Interval struct {
	stop chan struct{}
}

func NewInterval(f func(), interval time.Duration) *Interval {
	i := &Interval{
		stop: make(chan struct{}),
	}
	go func() {
		ticker := time.NewTicker(interval * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				f()
			case <-i.stop:
				return
			}
		}
	}()
	return i
}

func (i *Interval) Stop() {
	close(i.stop)
}
