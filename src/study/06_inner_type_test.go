package test

import (
	"fmt"
	"testing"
)

// 声明一个新的类型 ChineseBook，ChineseBook里头有一个 Book
// 同时ChineseBook 会自动继承Book 的所有方法
type ChineseBook struct {
	Book
	publisher string
}

// 测试内嵌类型
func TestInnerType(t *testing.T) {
	// 声明时可以看出来 ChineseBook 的接口是内嵌了一个 Book
	chineseBook := ChineseBook{Book{"chineseBook", 500.0, "编程语言chineseBook", "张三chineseBook", 100}, "出版社"}
	fmt.Println(chineseBook)
	// 两者是一样的
	fmt.Println("chineseBook.Name", chineseBook.Name, "| chineseBook.Book.Name", chineseBook.Book.Name)

	// Book 的方法，ChineseBook 可以直接调用
	chineseBook.salePrice()
	chineseBook.salePriceLocal("元")

	// Book 实现了 SalePrice 接口， ChineseBook 可以简介实现了 SalePrice 接口
	printSalePrice(chineseBook)
}

// 测试内嵌类型
func TestInnerType01(t *testing.T) {
	// 声明时可以看出来 ChineseBook 的接口是内嵌了一个 Book
	chineseBook := ChineseBook1{Book{"chineseBook", 500.0, "编程语言chineseBook", "张三chineseBook", 100}, "出版社"}
	fmt.Println(chineseBook)

	//如果不是语法糖的嵌入，方法继承和接口实现继承是无法实现的
	// Book 的方法，ChineseBook 可以直接调用
	//chineseBook.salePrice()
	//chineseBook.salePriceLocal("元")
	//
	//// Book 实现了 SalePrice 接口， ChineseBook 可以简介实现了 SalePrice 接口
	//printSalePrice(chineseBook)
}

type ChineseBook1 struct {
	MyBook    Book
	publisher string
}
