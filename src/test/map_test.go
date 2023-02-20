package test

import "testing"

func TestMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))
	delete(m1, 3)
	t.Logf("after delete 3 len m1=%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Log(m1[0])
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestExist(t *testing.T) {
	m1 := map[int]int{1: 1}
	t.Log(m1[1])
	t.Log(m1[2])
	if v, ok := m1[3]; ok {
		t.Logf("v=%d", v)
	} else {
		t.Log("v is not exit")
	}
}

func TestFor(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Logf("k=%d v=%d", k, v)
	}
}

func TestMapFunc(t *testing.T) {
	m := map[string]func(op int) int{}

	m["num"] = func(op int) int {
		return op
	}
	m["double"] = func(op int) int {
		return op * 2
	}
	m["treble"] = func(op int) int {
		return op * 3
	}

	t.Log("num", m["num"](2))
	t.Log("double", m["double"](2))
	t.Log("treble", m["treble"](2))
}
