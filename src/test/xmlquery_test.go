package test

import (
	"context"
	"github.com/antchfx/xmlquery"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"testing"
	"time"
)

func TestMongo(t *testing.T) {
	ctx := context.Background()
	// db.getCollection("detail_4K原版").find().limit(1000).skip(0)
	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://192.168.3.5:27017", Database: "spider1", Coll: "detail"})
	if err != nil {
		panic(err)
	}
	cli.InsertOne(ctx, Movie{})
	one := Movie{}
	find := cli.Find(ctx, bson.M{})
	t.Log(find.One(&one), one)
}

type Movie struct {
	MovieId    string    `bson:"movie_id"`
	Html       string    `bson:"html"`
	Url        string    `bson:"url"`
	Type       string    `bson:"type"`
	CreateDate time.Time `bson:"create_date"`
}

func TestXmlquery(t *testing.T) {
	s := `<?xml version="1.0" encoding="UTF-8"?>

<bookstore>

<book category="COOKING">
  <title lang="en">Everyday Italian</title>
  <author>Giada De Laurentiis</author>
  <year>2005</year>
  <price>30.00</price>
</book>

<book category="CHILDREN">
  <title lang="en">Harry Potter</title>
  <author>J K. Rowling</author>
  <year>2005</year>
  <price>29.99</price>
</book>

<book category="WEB">
  <title lang="en">XQuery Kick Start</title>
  <author>James McGovern</author>
  <author>Per Bothner</author>
  <author>Kurt Cagle</author>
  <author>James Linn</author>
  <author>Vaidyanathan Nagarajan</author>
  <year>2003</year>
  <price>49.99</price>
</book>

<book category="WEB">
  <title lang="en">Learning XML</title>
  <author>Erik T. Ray</author>
  <year>2003</year>
  <price>39.95</price>
</book>

</bookstore>`

	doc, err := xmlquery.Parse(strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	var arr = []AttrXpath{AttrXpath{"//title", "lang", "lang"}}
	var config = Config{"//book", arr}
	Parser(t, doc, config)

}

func Parser(t *testing.T, doc *xmlquery.Node, config Config) {
	nodes := xmlquery.Find(doc, config.NodeXpath)

	for index, node := range nodes {
		ctx := map[string]string{}
		for _, attrParser := range config.AttrXpathArr {
			if attrParser.xpath != "" {
				element := node.SelectElement(attrParser.xpath)
				ctx[attrParser.alias] = element.SelectAttr(attrParser.attName)
			} else {
				ctx[attrParser.alias] = node.SelectAttr(attrParser.attName)
			}
		}
		t.Log(index, ctx)
	}
}

type Config struct {
	NodeXpath    string
	AttrXpathArr []AttrXpath
}
type AttrXpath struct {
	xpath   string
	attName string
	alias   string
}
