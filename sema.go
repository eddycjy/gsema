package gsema

import "sync"

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
	s.wg.Add(delta)
	for i := 0; i < delta; i++ {
		s.c <- struct{}{}
	}
}

func (s *Semaphore) Done() {
	<-s.c
	s.wg.Done()
}

func (s *Semaphore) Wait() {
	s.wg.Wait()
}
