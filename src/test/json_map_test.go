package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/antchfx/htmlquery"
	"reflect"
	"strings"
	"testing"
)

const xpathResultHtmlType = "html"
const xpathResultContentType = "content"
const xpathResultAttrType = "attr"

func TestJsonMap(t *testing.T) {
	jsonBuf := `
    {
    "titleArr": ["//title"],
    "textArr_content": ["//title"],
    "content_attr": ["//meta"],
    "title": "//title",
    "text_content": "//title",
    "head": "//head",
    "headText_content": "//head",
    "subjects": [
        "Go",
        "C++",
        "Python",
        "Test"
    ],
	"obj" : {
		"qwe":"123"
    }
}`
	//创建一个map
	m := make(map[string]interface{}, 4)

	err := json.Unmarshal([]byte(jsonBuf), &m) //第二个参数要地址传递
	if err != nil {
		printlnErr(err, "Parser config err")
		return
	}
	resultData := make(map[string]interface{}, 4)
	var html = "<!DOCTYPE html> <html> <head> <meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" /> <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title>  <title>asdasdasdasdasdsa -  asdasdasdasdasdsa -  sjdlkasjdlasd!</title> </body> </html>"
	//类型断言, 值，它是value类型
	for key, value := range m {
		//fmt.Printf("%v ============> %v\n", key, value)
		switch data := value.(type) {
		case string:
			if value == "" {
				continue
			}
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
				return
			}
			element, err := htmlquery.Query(doc, data)
			if err != nil {
				printlnErr(err, "Parser "+key+": "+data+" err")
				return
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
		case []interface{}:
			stringArr, err := toStringArr(value)
			if err != nil {
				printlnErr(err, "Parser "+key+" err")
				return
			}
			dataValue := stringArr[0]
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
				return
			}
			elements, err := htmlquery.QueryAll(doc, dataValue)
			if err != nil {
				printlnErr(err, "Parser "+key+": "+dataValue+" err")
				return
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
		case interface{}:
			fmt.Printf("map[%s]的值类型为{}interface, value = %v\n", key, data)
		}
	}
	fmt.Println(resultData)
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
	bytes, _ := json.Marshal(Result{false, msg, err.Error()})
	fmt.Println(string(bytes))
}

type Result struct {
	Success bool
	Msg     string
	Data    string
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
