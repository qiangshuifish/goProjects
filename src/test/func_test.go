package test

import (
	"fmt"
	"testing"
	"time"
)

func TestFunc(t *testing.T) {
	ret := timeSpent(func(op int) int {
		time.Sleep(time.Second * 2)
		return op * 2
	})
	t.Log("4 * 2 = ", ret(4))
}

func TestVarParam(t *testing.T) {
	count := 0
	add := func(ops ...int) int {
		for _, v := range ops {
			count = count + v
		}
		return count
	}
	t.Log("add 1, 2, 3 = ", add(1, 2, 3))
}

func TestMultiReturn(t *testing.T) {
	code, msg := multiReturn()
	t.Log("code :", code, "msg:", msg)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		now := time.Now()
		ret := inner(n)
		fmt.Println("time spent: ", time.Since(now).Seconds())
		return ret
	}
}

func multiReturn() (int, string) {
	return 1, "success"
}
