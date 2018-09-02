package sleep

import (
	"sync"
	"time"
)

type Sleep struct {
	duration time.Duration
	mutex    *sync.Mutex
	count    int
	done     chan bool
}

func New(duration time.Duration) *Sleep {
	return &Sleep{
		duration: duration,
		mutex:    &sync.Mutex{},
		done:     make(chan bool, 10000),
	}
}

func (s *Sleep) Sleep() {
	time.Sleep(s.duration)
	s.mutex.Lock()
	s.count++
	s.mutex.Unlock()
	s.done <- true
}

func (s *Sleep) Count() int {
	return s.count
}

func (s *Sleep) Done() chan bool {
	return s.done
}
