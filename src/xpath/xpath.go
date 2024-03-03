package main

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/pflag"
	utils "spider/src/utils"
)

func main() {
	var xpath = pflag.StringP("exp", "e", "", "使用表达式解析")
	var config = pflag.StringP("config", "c", "", "使用配置解析")
	var html = pflag.StringP("html", "h", "", "需要解析的html内容")
	var enableBase64 = pflag.BoolP("enableBase64", "b", false, "是否开启对入参开启base64编码")
	var resultHandle = pflag.StringP("resultHandle", "r", "", "结果处理器")

	pflag.Parse()
	if *enableBase64 {
		decodeXpath, _ := base64.StdEncoding.DecodeString(*xpath)
		decodeHtml, _ := base64.StdEncoding.DecodeString(*html)
		decodeConfig, _ := base64.StdEncoding.DecodeString(*config)
		decodeHandel, _ := base64.StdEncoding.DecodeString(*resultHandle)
		*xpath = string(decodeXpath)
		*html = string(decodeHtml)
		*config = string(decodeConfig)
		*resultHandle = string(decodeHandel)
	}
	s := utils.Exe(*config, *xpath, *html, *resultHandle)
	fmt.Println(s)
}
