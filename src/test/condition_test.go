package test

import "testing"

func TestName(t *testing.T) {
	if a, err := someFun(); err == nil {
		t.Log(a == 1)
	} else {
		t.Log(a == 1)
	}
}
func someFun() (int, error) {
	return 1, nil
}

func TestSwitch(t *testing.T) {
	a := one
	switch a {
	case one:
		t.Log("one")
	case two:
		t.Log("two")
	case three, four:
		t.Log("three,four")
	}
}

func TestSwitch1(t *testing.T) {
	a := one
	switch {
	case a == one:
		t.Log("one")
	case a == two:
		t.Log("two")
	case a == three || a == four:
		t.Log("three,four")
	default:
		t.Log("default")
	}
}
