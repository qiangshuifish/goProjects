package test

import (
	"fmt"
	"testing"
)

func TestArr(t *testing.T) {
	// 标准定义格式
	var arr = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// 指定长度
	var arr1 = [5]int{1, 2}
	fmt.Println(arr1)

	// 指定长度,指定位置赋值
	var arr2 = [5]int{1: 10, 4: 20}
	fmt.Println(arr2)

	// 访问数组
	fmt.Println("arr2[4]", arr2[4])

	// 修改数组的值
	arr2[3] = 30
	fmt.Println("arr2[3]", arr2[3])
}

// test 指针类型的数组
func TestArrPoint(t *testing.T) {
	var arr = [5]*int{1: new(int), 2: new(int)}

	fmt.Println("地址 arr[1]", arr[1])
	fmt.Println("值  *arr[1]", *arr[1])
}

// 数组复制
func TestArrCopy(t *testing.T) {

	var arr = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	var arr2 = arr
	fmt.Println("arr2 = arr |", "arr:", arr, "arr2", arr2)

	// 这里只有数组 arr2 的第一个值发生了变化
	// arr 的值不会变化
	// Go语言的数组复制是复制一个数组，不是复制的指针
	arr2[0] = 2
	fmt.Println("arr[0] = 2 |", "arr:", arr, "arr2", arr2)

	// 这里把 arr 的指针赋值给了 arr3，所以修改 arr3 也会导致 arr 的值发生变化
	var arr3 = &arr
	arr3[0] = 2
	fmt.Println("arr[0] = 2 |", "arr:", arr, "arr3", arr3)
}

// 多维数组
func TestManyDimensionArr(t *testing.T) {
	// 定义一个多维数组，并赋值
	var arr = [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("arr", arr)

	// 指定位置赋值
	var arr2 = [2][2]int{1: {3, 4}}
	fmt.Println("arr2", arr2)

	// 多维数组访问
	fmt.Println("arr[0]:", arr[0])
	fmt.Println("arr[0][0]:", arr[0][0])

	arr[1] = [...]int{123, 456}
	fmt.Println("arr[1] = [...]int{123,456} :", arr[1])

	arr[1][1] = 100
	fmt.Println("arr[1][1] = 100 :", arr[1][1])
}
