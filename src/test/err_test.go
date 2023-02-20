package test

import (
	"errors"
	"testing"
)

func TestErr(t *testing.T) {
	code, err := returnErr()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(code)
	}
}

func returnErr() (int, error) {
	return 0, errors.New("something err")
}

func TestPanic(t *testing.T) {
	defer func() {
		t.Log("TestPanic defer exe")
	}()
	t.Log("TestPanic before panic")
	panic(func() {
		t.Log("TestPanic panic exe")
	})
	t.Log("TestPanic after panic")
}

func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("recover err:", err)
		}
	}()
	t.Log("TestPanic before panic")
	panic("something err")
	t.Log("TestPanic after panic")
}
