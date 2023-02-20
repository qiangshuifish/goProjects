package test

import "testing"

func TestOperator(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 4, 5}
	//c:=[...]int{1,2,3,4,5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c)  不同长度定数组无法比较
	t.Log(a == d)
}

// 按位置零
