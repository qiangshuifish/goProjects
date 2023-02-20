package test

import (
	"strconv"
	"strings"
	"testing"
)

func Test_string(t *testing.T) {
	str := "文化a1"
	t.Log("len(srt)", len(str))
	t.Log("srt[1]", str[1])
	strRunes := []rune(str)
	t.Log("strRunes", strRunes)
	t.Log("len(strRunes)", len(strRunes))
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	split := strings.Split(s, ",")
	for _, s := range split {
		t.Log(s)
	}
	t.Log(strings.Join(split, "-"))
}

func TestStringConv(t *testing.T) {
	str := strconv.Itoa(10)
	t.Log("str ", str)

	if atoi, err := strconv.Atoi("10"); err == nil {
		t.Log("atoi", atoi)
	}
}
