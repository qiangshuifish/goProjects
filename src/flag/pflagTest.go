package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

// 用法1： 使用指针获取传入参数  --age 122 ，
var age *int = flag.Int("age", 1234, "help message for age")

//用法2： 短参数用法，传入参数 --age1 122 或者  -a 133

var age1 = flag.IntP("age1", "a", 222, "help message")

// 用法3： 定义一个变量，使用 flag.IntVar
var flagvar int
var enable *bool = flag.Bool("enable", true, "help message for enable")

func main() {
	// 定义了所有的flag之后调用一次 flag.Parse() 对应的flag能够获得值
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
	flag.Parse()
	fmt.Println("age:", *age)
	fmt.Println("age1", *age1)
	fmt.Println("enable:", *enable)

}
