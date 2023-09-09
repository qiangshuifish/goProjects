package test

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {
	var goLang = Book{}
	goLang = Book{"GoLang", 50.0, "编程语言", "张三", 0}
	fmt.Println(goLang)
	fmt.Println(goLang.Name)

	goLang.Name = "GO语言"
	goLang.id = 123
	fmt.Println(goLang)
	fmt.Println(&goLang)
	fmt.Println(*&goLang)

	var goLangPoint *Book = &goLang
	goLangPoint.Author = "李四"
	fmt.Println(goLangPoint)
}

// 结构体增加方法，结构作为参数使用
func TestStructMethod(t *testing.T) {
	goLang := Book{"GoLang", 50.0, "编程语言", "张三", 0}
	// 给结构体增加一个方法   func (book Book) info() {}
	fmt.Println("before goLang.info()", goLang)
	goLang.info()
	fmt.Println("after goLang.info()", goLang)

	fmt.Println()

	// 结构体作为方法参数   func printBook(book Book) {}
	fmt.Println("before printBook(goLang)", goLang)
	printBook(goLang)
	fmt.Println("after printBook(goLang)", goLang)
	fmt.Println()
}

// 如果类型是指针，传递引用会改变原来结构体的值
func TestStructPointer(t *testing.T) {
	goLang := Book{"GoLang", 50.0, "编程语言", "张三", 0}
	// 给结构体增加一个方法   func (book Book) info() {}
	fmt.Println("before goLang.info()", goLang)
	goLang.infoPointer()
	fmt.Println("after goLang.info()", goLang)

	fmt.Println()

	// 结构体作为方法参数   func printBook(book Book) {}
	fmt.Println("before printBook(goLang)", goLang)
	printBookWithPoint(&goLang)
	fmt.Println("after printBook(goLang)", goLang)
	fmt.Println()

}

func (book *Book) infoPointer() {
	book.Name = "notGoLang"
	fmt.Println(book)
}

// 方法之间传递的也是zhi，不是引用
func printBookWithPoint(book *Book) {
	book.id = 10000
	fmt.Println(book)
}

// 方法之间传递的也是zhi，不是引用
func printBook(book Book) {
	book.id = 10000
	fmt.Println(book)
}

func (book Book) info() {
	book.Name = "notGoLang"
	fmt.Println(book)
}

type Book struct {
	Name    string
	Price   float32
	Subject string
	Author  string
	id      int
}

// 已Book声明 EnglishBook，但是仍然是不同的类型
type EnglishBook Book
