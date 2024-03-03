package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xmlquery"
	"golang.org/x/net/html"
	"reflect"
	"strings"
)

const xpathResultHtmlType = "html"
const xpathResultContentType = "content"
const xpathResultAttrType = "attr"

func Exe(config string, xpath string, html string, resultHandle string) string {
	if config == "" && xpath == "" {
		printlnErr(errors.New("config or xpath must have one value"), "config or xpath must have one value")
		return "{}"
	}
	if config == "" {
		resultData := make(map[string]interface{}, 1)
		parser("name", html, xpath, resultData, true, false)
		elms := resultData["name"].([]string)

		// 如果需要处理结果参数
		if config == "" && resultHandle != "" && xpath != "" && len(elms) > 0 {
			result, _ := handleResult(elms, resultHandle)
			if result != nil {
				return printlnResult(result)
			}
		} else {
			// 到这里是不需要解析
			return printlnResult(elms)
		}
		return "{}"
	}

	configJson := make(map[string]interface{}, 4)
	err := json.Unmarshal([]byte(config), &configJson) //第二个参数要地址传递
	if err != nil {
		printlnErr(err, "Parser config err")
		return "{}"
	}
	if len(configJson) <= 0 {
		printlnErr(err, "Parser config err ,config is empty")
		return "{}"
	}
	resultData := make(map[string]interface{}, len(configJson))
	if xpathParser(configJson, html, resultData, false) {
		return "{}"
	}
	return printlnResult(resultData)
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
		xpathParser(configJson, elms[index], parserData, true)
		resultData[index] = parserData
	}

	// 返回结果
	return resultData, nil
}

func xpathParser(config map[string]interface{}, html string, resultData map[string]interface{}, useXml bool) bool {
	for key, value := range config {
		//fmt.Printf("%v ============> %v\n", key, value)
		switch data := value.(type) {
		case string:
			if value == "" {
				continue
			}
			parser(key, html, data, resultData, false, useXml)
		case []interface{}:
			stringArr, err := toStringArr(value)
			if err != nil {
				printlnErr(err, "Parser "+key+" err")
				return true
			}
			parser(key, html, stringArr[0], resultData, true, useXml)
		case interface{}:
			var subConfig = data.(map[string]interface{})
			subResultData := make(map[string]interface{}, 4)
			xpathParser(subConfig, html, subResultData, useXml)
			resultData[key] = subResultData
		}
	}
	return false
}
func parser(key string, html string, xpath string, resultData map[string]interface{}, arrMode bool, useXml bool) {
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
	if useXml {
		xmlQuery(key, html, xpath, resultData, arrMode, nameType, name)
	} else {
		htmlQuery(key, html, xpath, resultData, arrMode, nameType, name)
	}

}

func htmlQuery(key string, html string, xpath string, resultData map[string]interface{}, arrMode bool, nameType string, name string) {
	doc, err := htmlquery.Parse(strings.NewReader(html))
	if err != nil {
		printlnErr(err, "Parser htmlxk err")
		return
	}
	if arrMode {
		elements, err := htmlquery.QueryAll(doc, xpath)
		if err != nil {
			printlnErr(err, "Parser "+key+": "+xpath+" err")
			return
		}
		arr := make([]string, len(elements))
		for index, element := range elements {
			arr[index] = getHtmlResultData(nameType, name, element)
		}
		resultData[name] = arr
	} else {
		element, err := htmlquery.Query(doc, xpath)
		if err != nil {
			printlnErr(err, "Parser "+key+": "+xpath+" err")
			return
		}
		resultData[name] = getHtmlResultData(nameType, name, element)
	}
}

func xmlQuery(key string, html string, xpath string, resultData map[string]interface{}, arrMode bool, nameType string, name string) {
	doc, err := xmlquery.Parse(strings.NewReader(html))
	if err != nil {
		printlnErr(err, "Parser htmlxk err")
		return
	}
	if arrMode {
		elements, err := xmlquery.QueryAll(doc, xpath)
		if err != nil {
			printlnErr(err, "Parser "+key+": "+xpath+" err")
			return
		}
		arr := make([]string, len(elements))
		for index, element := range elements {
			arr[index] = getXmlResultData(nameType, name, element)
		}
		resultData[name] = arr
	} else {
		element, err := xmlquery.Query(doc, xpath)
		if err != nil {
			printlnErr(err, "Parser "+key+": "+xpath+" err")
			return
		}
		resultData[name] = getXmlResultData(nameType, name, element)
	}
}

func getHtmlResultData(nameType string, name string, element *html.Node) string {
	var data = ""
	defer func() {
		recover()
	}()
	if nameType == xpathResultHtmlType {
		data = htmlquery.OutputHTML(element, true)
	}
	if nameType == xpathResultContentType {
		data = strings.TrimSpace(htmlquery.InnerText(element))
	}
	if nameType == xpathResultAttrType {
		data = strings.TrimSpace(htmlquery.SelectAttr(element, name))
	}
	return data
}

func getXmlResultData(nameType string, name string, element *xmlquery.Node) string {
	var data = ""
	defer func() {
		recover()
	}()
	if nameType == xpathResultHtmlType {
		data = element.OutputXML(true)
	}
	if nameType == xpathResultContentType {
		data = strings.TrimSpace(element.InnerText())
	}
	if nameType == xpathResultAttrType {
		data = strings.TrimSpace(element.SelectAttr(name))
	}
	return data
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

func printlnResult(marshal interface{}) string {
	bytes, _ := json.Marshal(Result{true, "", marshal})
	result := string(bytes)
	result = strings.Replace(result, "\\u003c", "<", -1)
	result = strings.Replace(result, "\\u003e", ">", -1)
	result = strings.Replace(result, "\\u0026", "&", -1)
	return result
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
