package test

import (
	"math"
	"testing"
)

func TestImplicit(t *testing.T) {
	var a int32 = 0
	var b int64
	b = int64(a)
	t.Log(b)
}

func TestMath(t *testing.T) {
	t.Log(math.MaxFloat64, math.MaxFloat32)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T,%T", a, aPtr)
}

func TestString(t *testing.T) {
	var str string
	var a int32
	t.Log(str, a)
}
