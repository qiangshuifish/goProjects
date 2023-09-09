package collector

import (
	"fmt"
	"github.com/gocolly/colly"
)

var retries = make(map[string]int)

func Retry(c *colly.Collector, retryTimes int) {
	if retryTimes == 0 {
		retryTimes = 3
	}
	// 请求失败时的处理
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("Request URL:", response.Request.URL, "failed with response:", response, "\nError:", err)
		if retries[response.Request.URL.String()] < retryTimes {
			retries[response.Request.URL.String()]++
			c.Visit(response.Request.URL.String())
		}
	})
}
