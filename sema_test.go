package gsema

import (
	"fmt"
	"testing"
	"time"
)

func TestSemaphore_Add(t *testing.T) {
	sema := NewSemaphore(10)
	defer sema.Close()
	workerTotal := 100
	for i := 0; i < workerTotal; i++ {
		go func() {
			defer sema.Add(-1)
			sema.Add(1)
			fmt.Printf("go func: %d, time: %s\n", i, time.Now().Format("2006-01-02 15:04:05"))
			time.Sleep(time.Second)
		}()
	}
	sema.Wait()
}

func TestSemaphore(t *testing.T) {
	sema := NewSemaphore(10)
	defer sema.Close()
	workerTotal := 100
	for i := 0; i < workerTotal; i++ {
		go func() {
			defer sema.Done()
			sema.Add(1)
			fmt.Printf("go func: %d, time: %s\n", i, time.Now().Format("2006-01-02 15:04:05"))
			time.Sleep(time.Second)
		}()
	}
	sema.Wait()
}

func BenchmarkSemaphore(b *testing.B) {
	sema := NewSemaphore(10)
	defer sema.Close()
	workerTotal := 100
	for i := 0; i < workerTotal; i++ {
		go func() {
			defer sema.Done()
			sema.Add(1)
			fmt.Printf("go func: %d, time: %s\n", i, time.Now().Format("2006-01-02 15:04:05"))
			time.Sleep(time.Second)
		}()
	}
	sema.Wait()
}
