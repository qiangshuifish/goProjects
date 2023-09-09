package test

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	// 创建一个长度和容量都是5的切片
	var slice = make([]int, 5)
	fmt.Println(slice)

	// 创建一个长度为3，容量为5的切片
	var slice1 = make([]int, 3, 5)
	//var arr1 = make([]int,5,3) 创建一个长度大于容量的会报错
	fmt.Println(slice1)

	slice2 := []int{20, 30, 40} // 和声明数组的区别就是[]里带不带参数 slice = [...]int{1,2,3,4,5}
	fmt.Println("slice", slice2)

	// 声明一个长度未10的切片，并给第10个赋值100
	slice3 := []int{9: 100}
	fmt.Println("slice3", slice3)

	// 定义空切片，空数组
	var nilSlice []int
	fmt.Println("nilSlice", nilSlice)
	fmt.Println("cap(nilSlice)", cap(nilSlice))
	fmt.Println("len(nilSlice)", len(nilSlice))
}

func TestUseSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("slice", slice)

	// 把slice的第2个到5个元素赋值给neSlice
	newSlice := slice[1:4]
	fmt.Println("slice", slice, "|", "newSlice", newSlice)
	// 定义切片 slice[i:j],假设 slice的容量是 k
	// 容量 cap = k - i
	// 长度 len = j - i
	fmt.Println("cap(newSlice)", cap(newSlice))
	fmt.Println("len(newSlice)", len(newSlice))

	// 切片底层是共用一个数组，改变 newSlice 的第1个元素，同时也会改变 slice 的第2个元素
	fmt.Println("=========================")
	fmt.Println("newSlice[0] = 100")
	newSlice[0] = 100
	fmt.Println("newSlice[0]", newSlice[0])
	fmt.Println("slice[1]", slice[1])
	fmt.Println("slice", slice, "|", "newSlice", newSlice)
}

func TestSliceAppend(t *testing.T) {
	slice := make([]int, 3, 5)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	fmt.Println(slice)

	fmt.Println("")

	// append 的元素在原来的容量范围内时，直接添加
	newSlice := append(slice, 4)
	fmt.Println("slice", slice, "newSlice", newSlice)
	fmt.Println("cap(slice)", cap(slice), "len(slice)", len(slice))
	fmt.Println("cap(newSlice)", cap(newSlice), "len(newSlice)", len(newSlice))
	slice[0] = 100
	fmt.Println("slice[0] = 100", "slice", slice, "newSlice", newSlice)

	fmt.Println("")

	// append 的元素在原来的容量范围外时，会创建一个新数组，这个数组会是原来的两倍，并将原来的值复制进去并追加新值
	// 容量小于1000时，每次扩容1背，大于1000时扩容1.25倍
	newSlice = append(slice, 4, 5, 6, 7)
	fmt.Println("slice", slice, "newSlice", newSlice)
	// 原来的 slice 还是使用的原来的数组
	fmt.Println("cap(slice)", cap(slice), "len(slice)", len(slice))
	// newSlice 已经使用了新数组
	fmt.Println("cap(newSlice)", cap(newSlice), "len(newSlice)", len(newSlice))

	// 所以这里不会改变新数组的值
	slice[0] = 1000
	fmt.Println("slice[0] = 1000", "slice", slice, "newSlice", newSlice)
}

// 切片操作语法糖
func TestSliceIndex(t *testing.T) {

	slice := []int{1, 2, 3, 4, 5}
	// newSlice 取slice第0位开始，向后取4位，并限制容量为5
	newSlice := slice[0:4:5]
	fmt.Println("slice", slice, "slice[1:3:2]", newSlice)
	fmt.Println("cap(newSlice)", cap(newSlice), "len(newSlice)", len(newSlice))
}

// 切片复制，切片复制的是引用，所以共享底层数据
func TestSliceCopy(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	newSlice := slice
	fmt.Println("slice", slice, "newSlice", newSlice)

	newSlice[0] = 100
	fmt.Println("newSlice[0] = 100", "slice", slice, "newSlice", newSlice)
}

func TestSliceFor(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	for i := range slice {
		fmt.Println("slice[i]", slice[i])
	}

	for i, v := range slice {
		fmt.Println("slice[i]", slice[i], "v", v)
	}
}
