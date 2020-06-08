package gsema

import (
	"sync"
)

type Semaphore struct {
	c  chan struct{}
	wg *sync.WaitGroup
}

func NewSemaphore(maxSize int) *Semaphore {
	return &Semaphore{
		c:  make(chan struct{}, maxSize),
		wg: new(sync.WaitGroup),
	}
}

func (s *Semaphore) Add(delta int) {
	if delta < 0 {
		for i := delta; i < 0; i++ {
			s.Done()
		}
	} else if delta > 0 {
		for i := 0; i < delta; i++ {
			s.c <- struct{}{}
			s.wg.Add(1)
		}
	}
}

func (s *Semaphore) Done() {
	<-s.c
	s.wg.Done()
}

func (s *Semaphore) Wait() {
	s.wg.Wait()
}

func (s *Semaphore) Close() {
	if len(s.c) == 0 {
		close(s.c)
	}
}
