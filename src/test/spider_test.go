package test

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"strings"
	"testing"
	"time"
)

func TestSpider(t *testing.T) {
	c := colly.NewCollector()
	c.OnResponse(func(res *colly.Response) {
		fmt.Println(res.Request.URL.String())
		html := string(res.Body)

		doc, err := htmlquery.Parse(strings.NewReader(html))
		if err != nil {
			panic(err)
		}
		elements, err := htmlquery.QueryAll(doc, "/html/body[@id='nv_forum']/div[@id='wp']/div[@class='boardnav']/div[@id='ct']/div[@class='mn']/div[@id='threadlist']/div[@class='bm_c']/form[@id='moderate']/table[@id='threadlisttableid']/tbody[@id]/tr/th[@class='new']/a[@class='s xst']")
		for _, element := range elements {
			fmt.Println("===========", htmlquery.OutputHTML(element, false))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.Visit("https://www.5sjx4.cn/forum-2-1.html")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
