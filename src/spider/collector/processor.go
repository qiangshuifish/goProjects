package collector

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"spider/src/utils"
	"strconv"
	"time"
)

type Config struct {
	RequestTimeOut time.Duration
	RetryTimes     int
	Express        string
	ResultHandle   string
	Config         string
	StartUrl       string
	ExeType        ExeType
}

type ExeType string

const (
	SimplePrint ExeType = "print"
	Save2File   ExeType = "save_to_file"
	Save2Mongo  ExeType = "save_to_mongodb"
)

var collection *mongo.Collection = nil

func Processor(config Config) func(response *colly.Response) {
	switch config.ExeType {
	case Save2File:
		return SaveToFile(config)
	case Save2Mongo:
		return SaveToMongo(config)
	case SimplePrint:
	default:
		return Print(config)
	}
	return func(response *colly.Response) {
	}
}

// Print 简单打印结果
func Print(config Config) (f colly.ResponseCallback) {
	return func(response *colly.Response) {
		fmt.Println("success", response.Request.URL.String(), getResult(response, config))
	}
}

// SaveToFile 保存结果到文件中
func SaveToFile(config Config) (f colly.ResponseCallback) {
	return func(response *colly.Response) {
		os.WriteFile(strconv.Itoa(int(response.Request.ID))+"html", []byte(getResult(response, config)), 0644)
	}
}

func getResult(response *colly.Response, config Config) string {
	exe := utils.Exe(config.Config, config.Express, string(response.Body), config.ResultHandle)
	return exe
}

// SaveToMongo 保存结果到 mongodb 中
func SaveToMongo(config Config) (f colly.ResponseCallback) {
	if collection == nil {
		panic(errors.New("collection not init"))
	}
	return func(response *colly.Response) {
		//result := Result{ "","","","", string(response.Body),time.Now()}
		var result = make(map[string]any)
		err := json.Unmarshal([]byte(getResult(response, config)), &result)
		result["html"] = string(response.Body)
		result["create_time"] = time.Now()
		result["url"] = response.Request.URL.String()
		insertResult, err := collection.InsertOne(context.TODO(), result)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(insertResult)
		}
	}
}

type Result struct {
	Name       string    `bson:"name,omitempty"`
	Type       string    `bson:"type,omitempty"`
	Url        string    `bson:"url,omitempty"`
	MoveId     string    `bson:"move_id,omitempty"`
	Html       string    `bson:"result,omitempty"`
	CreateTime time.Time `bson:"create_time,omitempty"`
}
