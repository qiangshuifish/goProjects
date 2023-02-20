package test

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("init exe")
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("finally  exe")
	}()

	t.Logf("TestDefer exe")
	panic("err")

	t.Logf("afetr panic")
}
