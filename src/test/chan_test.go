package test

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	asyncRet := AsyncService()
	other()
	t.Log(<-asyncRet)
}

func AsyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("service return:", ret)
		retCh <- ret
		fmt.Println("service Done:")
	}()
	return retCh
}

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "service success"
}
func other() {
	time.Sleep(time.Millisecond * 70)
	fmt.Println("other success")
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan string, 0)
	for i := 0; i < 5; i++ {
		go func(i int) {
			for {
				if isCancelled(cancelChan) {
					t.Log(time.Now().Format("2006-01-02"), "chan id :", i, " cancel")
					break
				}
				t.Log(time.Now().Format("2006-01-02"), "chan id :", i, " isRun")
				time.Sleep(time.Millisecond * 10)
			}
		}(i)
	}
	t.Log(time.Now().Format("2006-01-02 15:04:05"), "main", "waite")
	time.Sleep(time.Millisecond * 100)
	//cancelChan <- "str" 这种方式只能关闭一个
	CancelChan(cancelChan)
	t.Log(time.Now().Format("2006-01-02 15:04:05"), "main to cancel")
	time.Sleep(time.Millisecond * 100)
	t.Log(time.Now().Format("2006-01-02 15:04:05"), "main to exit")
}

func CancelChan(cancelChan chan string) {
	close(cancelChan)
}

func isCancelled(cancelChan chan string) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}
