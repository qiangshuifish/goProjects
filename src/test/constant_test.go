package test

import "testing"

const (
	one = iota + 1
	two
	three
	four
)

func TestConst(t *testing.T) {
	t.Log(one, two, three, four)
}

const (
	open = iota << 1
	closed
	Pending
)

func TestByte(t *testing.T) {
	t.Log(open, closed, Pending)
}

func TestVar(t *testing.T) {
	var a int = 0
	b := 1
	var (
		c = 2
		d = 3
	)
	t.Log(a, b, c, d)

	a, b = b, a
	t.Log(a, b, c, d)
}
