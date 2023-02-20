package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/spf13/pflag"
	"reflect"
	"strings"
)

const xpathResultHtmlType = "html"
const xpathResultContentType = "content"
const xpathResultAttrType = "attr"

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
	exe(*config, *xpath, *html, *resultHandle)
}
func exe(config string, xpath string, html string, resultHandle string) {
	if config == "" && xpath == "" {
		printlnErr(errors.New("config or xpath must have one value"), "config or xpath must have one value")
		return
	}
	if config == "" {
		resultData := make(map[string]interface{}, 1)
		parser("name", html, xpath, resultData, true)
		elms := resultData["name"].([]string)

		// 如果需要处理结果参数
		if config == "" && resultHandle != "" && xpath != "" && len(elms) > 0 {
			result, _ := handleResult(elms, resultHandle)
			if result != nil {
				printlnResult(result)
			}
		} else {
			// 到这里是不需要解析
			printlnResult(elms)
		}
		return
	}

	configJson := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(config), &configJson) //第二个参数要地址传递
	if err != nil {
		printlnErr(err, "Parser config err")
		return
	}
	if len(configJson) <= 0 {
		printlnErr(err, "Parser config err ,config is empty")
		return
	}
	resultData := make(map[string]interface{}, len(configJson))
	if xpathParser(configJson, html, resultData) {
		return
	}
	printlnResult(resultData)
}

func handleResult(elms []string, resultHandle string) ([]map[string]interface{}, error) {
	// 解析配置
	configJson := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(resultHandle), &configJson)
	if err != nil {
		printlnErr(err, "parser resultHandle data err:"+resultHandle)
		return nil, err
	}

	// 解析 html 标签
	resultData := make([]map[string]interface{}, len(elms))
	for index := range elms {
		parserData := make(map[string]interface{}, len(configJson))
		xpathParser(configJson, elms[index], parserData)
		resultData[index] = parserData
	}

	// 返回结果
	return resultData, nil
}

func xpathParser(config map[string]interface{}, html string, resultData map[string]interface{}) bool {
	for key, value := range config {
		//fmt.Printf("%v ============> %v\n", key, value)
		switch data := value.(type) {
		case string:
			if value == "" {
				continue
			}
			parser(key, html, data, resultData, false)
		case []interface{}:
			stringArr, err := toStringArr(value)
			if err != nil {
				printlnErr(err, "Parser "+key+" err")
				return true
			}
			parser(key, html, stringArr[0], resultData, true)
		case interface{}:
			var subConfig = data.(map[string]interface{})
			subResultData := make(map[string]interface{}, 4)
			xpathParser(subConfig, html, subResultData)
			resultData[key] = subResultData
		}
	}
	return false
}
func parser(key string, html string, xpath string, resultData map[string]interface{}, arrMode bool) bool {
	var name string
	var nameType = xpathResultHtmlType
	if strings.Contains(key, "_") {
		split := strings.Split(key, "_")
		name = split[0]
		if len(split) <= 1 || split[1] == "" {
			nameType = xpathResultHtmlType
		}
		nameType = getXpathResultType(split[1])
	} else {
		name = key
	}
	doc, err := htmlquery.Parse(strings.NewReader(html))
	if err != nil {
		printlnErr(err, "Parser html err")
		return false
	}
	if arrMode {
		elements, err := htmlquery.QueryAll(doc, xpath)
		if err != nil {
			printlnErr(err, "Parser "+key+": "+xpath+" err")
			return false
		}
		arr := make([]string, len(elements))
		for index, element := range elements {
			if nameType == xpathResultHtmlType {
				arr[index] = htmlquery.OutputHTML(element, true)
			}
			if nameType == xpathResultContentType {
				arr[index] = strings.TrimSpace(htmlquery.InnerText(element))
			}
			if nameType == xpathResultAttrType {
				arr[index] = strings.TrimSpace(htmlquery.SelectAttr(element, name))
			}
		}
		resultData[name] = arr
	} else {
		element, err := htmlquery.Query(doc, xpath)
		if err != nil {
			printlnErr(err, "Parser "+key+": "+xpath+" err")
			return false
		}
		if nameType == xpathResultHtmlType {
			resultData[name] = htmlquery.OutputHTML(element, true)
		}
		if nameType == xpathResultContentType {
			resultData[name] = strings.TrimSpace(htmlquery.InnerText(element))
		}
		if nameType == xpathResultAttrType {
			resultData[name] = strings.TrimSpace(htmlquery.SelectAttr(element, name))
		}
	}
	return true
}
func getXpathResultType(split string) string {
	switch split {
	case xpathResultHtmlType:
		return xpathResultHtmlType
	case xpathResultContentType:
		return xpathResultContentType
	case xpathResultAttrType:
		return xpathResultAttrType
	}
	return xpathResultHtmlType
}
func printlnErr(err error, msg string) {
	bytes, _ := json.Marshal(Result{false, msg, err})
	fmt.Println(string(bytes))
}

func printlnResult(marshal interface{}) {
	bytes, _ := json.Marshal(Result{true, "", marshal})
	result := string(bytes)
	result = strings.Replace(result, "\\u003c", "<", -1)
	result = strings.Replace(result, "\\u003e", ">", -1)
	result = strings.Replace(result, "\\u0026", "&", -1)
	fmt.Println(result)
}

type Result struct {
	Success bool
	Msg     string
	Data    interface{}
}

func toStringArr(actual interface{}) ([]string, error) {
	var res []string
	value := reflect.ValueOf(actual)
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		return nil, errors.New("parse error")
	}
	for i := 0; i < value.Len(); i++ {
		res = append(res, value.Index(i).Interface().(string))
	}
	return res, nil
}
