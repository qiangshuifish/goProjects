package test

import (
	"sync"
	"testing"
	"time"
)

func TestGroutineTest(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			t.Log(i)
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
}

func TestThreadSafe(t *testing.T) {
	var mut sync.Mutex
	count := 0
	for i := 0; i < 5000; i++ {
		go func(i int) {
			count++
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
	t.Log("count:", count)

	countSate := 0
	for i := 0; i < 5000; i++ {
		go func(i int) {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			countSate++
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
	t.Log("countSate:", countSate)

}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var mut sync.Mutex
	countSate := 0
	for i := 0; i < 5000; i++ {
		go func(i int) {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			countSate++
			wg.Done()
		}(i)
	}
	wg.Wait()
	t.Log("countSate:", countSate)
}
