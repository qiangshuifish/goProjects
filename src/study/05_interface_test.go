package test

import (
	"fmt"
	"testing"
)

// 声明一个接口
type SalePrice interface {
	salePrice()

	salePriceLocal(unit string)
}

// 让 Book 实现接口
func (b Book) salePrice() {
	fmt.Println("价格为：", b.Price)
}

// 绑定带参数的接口
func (b Book) salePriceLocal(unit string) {
	fmt.Println("价格为：", b.Price, unit)
}

// 通过绑定的方式，让 struct 继承接口
func TestInterface(t *testing.T) {
	// 声明了一个book
	goLang := Book{"GoLang", 50.0, "编程语言", "张三", 0}

	// 把 Book 当做 SalePrice 传递了进去
	printSalePrice(goLang)
}

// 多态方法，必须要实现了 SalePrice 所有的方法的结构体才能传递进来
func printSalePrice(book SalePrice) {
	book.salePrice()
	book.salePriceLocal("元")
}

// 多态测试
func TestPolymorphic(t *testing.T) {
	// 声明了一个book
	goLang := Book{"GoLang", 50.0, "编程语言", "张三", 0}
	chinese := EnglishBook{"EnglishBook", 500.0, "编程语言EnglishBook", "张三EnglishBook", 100}

	// 把 Book 当做 SalePrice 传递了进去
	printSalePrice(goLang)

	printSalePrice(chinese)
}

// 让 Book 实现接口
func (b EnglishBook) salePrice() {
	fmt.Println("ChineseBook价格为：", b.Price)
}

// 绑定带参数的接口
func (b EnglishBook) salePriceLocal(unit string) {
	fmt.Println("ChineseBook价格为：", b.Price, unit)
}
