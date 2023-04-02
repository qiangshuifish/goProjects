package main

import (
	"bufio"
	"fmt"
	flag "github.com/spf13/pflag"
	"log"
	"os"
)

// 用法1： 使用指针获取传入参数  --age 122 ，
var age *int = flag.Int("age", 1234, "help message for age")

//用法2： 短参数用法，传入参数 --age1 122 或者  -a 133

var age1 = flag.IntP("age1", "a", 222, "help message")

// 用法3： 定义一个变量，使用 flag.IntVar
var flagvar int
var enable *bool = flag.Bool("enable", true, "help message for enable")
var str = flag.String("str", "", "help message for enable")

func main() {
	// 定义了所有的flag之后调用一次 flag.Parse() 对应的flag能够获得值
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
	flag.Parse()

	args := flag.Args()

	fileInfo, _ := os.Stdin.Stat()
	if (fileInfo.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
		log.Fatal("The command is intended to work with pipes.")
	}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println("pipeline", s.Text())
	}

	fmt.Println("age:", *age)
	fmt.Println("age1", *age1)
	fmt.Println("enable:", *enable)
	fmt.Println("str:", *str)
	fmt.Println("args:", args)
}
