package main

import (
	"github.com/go-redis/redis"
	"github.com/gocolly/colly"
	"net/url"
	"spider/src/spider/collector"
	"strconv"
	"time"
)

func main() {
	main01()
}

func main01() {
	config := collector.Config{}
	config.RequestTimeOut = 500 * time.Millisecond
	config.Express = "/html/body/div[@id='wrapper']/div[@id='content']/div[@class='grid-16-8 clearfix']/div[@class='aside']/ul[@class='hot-tags-col5 s']/li[1]/ul[@class='clearfix']/li/a[@class='tag']"
	config.ResultHandle = `{"value_content":"//a","href_attr":"//a"}`
	config.Config = ""
	config.StartUrl = "https://book.douban.com/"

	c := colly.NewCollector(colly.AllowURLRevisit(), colly.IgnoreRobotsTxt())

	c.SetRequestTimeout(config.RequestTimeOut)
	collector.Retry(c, config.RetryTimes)

	// 查找并访问所有链接
	c.OnResponse(collector.Processor(config))

	c.Visit(config.StartUrl)
	c.Wait()
}

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(addr string, password string, db int) *RedisStorage {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisStorage{
		client: client,
	}
}

func (r *RedisStorage) Init() error {
	return nil
}

func (r *RedisStorage) Visited(requestID uint64) error {
	return r.client.Set(strconv.FormatUint(requestID, 10), "true", 0).Err()
}

func (r *RedisStorage) IsVisited(requestID uint64) (bool, error) {
	result, e := r.client.Exists(strconv.FormatUint(requestID, 10)).Result()
	return result == 1, e
}

func (r *RedisStorage) Cookies(u *url.URL) string {
	cookies, err := r.client.Get(u.String()).Result()
	if err != nil {
		return ""
	}
	return cookies
}

func (r *RedisStorage) SetCookies(u *url.URL, cookies string) {
	r.client.Set(u.String(), cookies, 0)
}
