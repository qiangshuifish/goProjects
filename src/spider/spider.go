package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/spf13/pflag"
)

var url = pflag.StringP("url", "u", "", "需要抓去的url")

func main() {
	pflag.Parse()
	if *url == "" {
		return
	}
	c := colly.NewCollector()
	c.OnResponse(func(response *colly.Response) {
		fmt.Println(string(response.Body))
	})
	c.Visit(*url)
}
