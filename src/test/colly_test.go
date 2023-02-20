package test

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xpath"
	"github.com/gocolly/colly"
	"strings"
	"testing"
)

func TestColly(t *testing.T) {
	c := colly.NewCollector()
	//detail := spider.NewCollector()

	c.OnResponse(func(res *colly.Response) {
		fmt.Println(res.Request.URL.String())
		html := string(res.Body)
		fmt.Println(html)
		fmt.Println("===============================")

		//
		doc, err := htmlquery.Parse(strings.NewReader(html))
		if err != nil {
			return
		}
		all, err := htmlquery.QueryAll(doc, "//div[@class='paginator']//a//@href")
		if err != nil {
			return
		}
		for _, node := range all {
			fmt.Println(htmlquery.SelectAttr(node, "href"))
			//c.Visit(htmlquery.SelectAttr(node,"href"))
		}

		detailList, err := htmlquery.QueryAll(doc, "//div[@class='indent']//tr[@class='item']")
		if err != nil {
			return
		}
		for _, item := range detailList {
			fmt.Println("item", htmlquery.OutputHTML(item, false))

			a := htmlquery.FindOne(item, "//div[@class='pl2']//score")
			title := htmlquery.SelectAttr(a, "title")
			fmt.Println(title)
			name1 := htmlquery.SelectAttr(item, "//div[@class='pl2']//a[0]//@title")
			fmt.Println(name1)
			expr, _ := xpath.Compile("//div[@class='pl2']//a[0]//@title")
			title1 := expr.Evaluate(htmlquery.CreateXPathNavigator(item))
			fmt.Println(title1)
		}

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.Visit("https://book.douban.com/top250")
}

type Spider struct {
	name string
}

type SpiderNode struct {
	name    string
	parsers ElementParser
	action  Action
}

type ElementParser struct {
	Name  string
	Xpath string
	attrs []AttrParser
}

type GetAttrs interface {
	GetAttrs() []string
}

func (p *ElementParser) GetAttrs() []string {
	return []string{"", ""}
}

type AttrParser struct {
	name  string
	xpath string
}

type Action struct {
}
