package utils

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client = nil

const dbUrl = ""
const dbName = ""

func init() {
	_, err := GetMongoClient(dbUrl)
	if err != nil {
		panic(err)
	}
}

// GetMongoClient 初始化配置
func GetMongoClient(url string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	return client, err
}

func GetDbByName(dbname string) (*mongo.Database, error) {
	if client == nil {
		return nil, errors.New("client not init ")
	}
	database := client.Database(dbname)
	return database, nil
}

func GetDb() (*mongo.Database, error) {
	return GetDbByName(dbName)
}

func GetCollection(collectionName string) (*CollectionClient, error) {
	return GetCollectionClient(dbName, collectionName)
}

func GetCollectionClient(dbname string, collectionName string) (*CollectionClient, error) {
	if client == nil {
		return nil, errors.New("client not init ")
	}
	collection := client.Database(dbname).Collection(collectionName)
	collectionClient := CollectionClient{collectionName, collection}
	return &collectionClient, nil
}

// CollectionClient 是操作 collection 的客户端
type CollectionClient struct {
	CollectionName string
	client         *mongo.Collection
}

// InsertOne 是 CollectionClient 结构体的一个方法。
// 这个方法接收一个 data 参数，表示要插入的文档。
// 这个方法的主要功能是在 MongoDB 数据库中执行单个文档的插入操作。
// 如果插入操作成功，则返回 true；如果插入操作失败，则返回 false。
func (c *CollectionClient) InsertOne(data interface{}) bool {
	// 使用 InsertOne 方法在 MongoDB 中执行单个文档的插入操作
	_, err := c.client.InsertOne(context.TODO(), data)
	// 如果在执行插入操作时发生错误，返回 false
	if err != nil {
		return false
	}
	// 插入操作成功，返回 true
	return true
}

// InsertMany 是 CollectionClient 结构体的一个方法。
// 这个方法接收一个可变参数 data，表示要插入的多个文档。
// 这个方法的主要功能是在 MongoDB 数据库中执行批量插入操作。
// 如果插入操作成功，则返回 true；如果插入操作失败，则返回 false。
func (c *CollectionClient) InsertMany(data ...interface{}) bool {
	// 使用 InsertMany 方法在 MongoDB 中执行批量插入操作
	_, err := c.client.InsertMany(context.TODO(), data)
	// 如果在执行插入操作时发生错误，返回 false
	if err != nil {
		return false
	}
	// 插入操作成功，返回 true
	return true
}

// QueryOne 是 CollectionClient 结构体的一个方法。
// 这个方法接收一个 bson.D 类型的查询参数，返回一个字符串和一个错误信息。
// 这个方法的主要功能是在 MongoDB 数据库中执行查询，并将查询结果转换为字符串返回。
func (c *CollectionClient) QueryOne(query bson.D) (string, error) {
	// 使用 FindOne 方法在 MongoDB 中执行查询，返回一个单一的结果
	one := c.client.FindOne(context.TODO(), query)
	// 将查询结果解码为字节
	bytes, err := one.DecodeBytes()
	// 如果在解码查询结果时发生错误，返回空字符串和错误信息
	if err != nil {
		return "", err
	}
	// 将查询结果的字节转换为字符串并返回
	return string(bytes), nil
}

// QueryMany 是 CollectionClient 结构体的一个方法。
// 这个方法接收一个 bson.D 类型的查询参数，返回一个字符串数组和一个错误信息。
// 这个方法的主要功能是在 MongoDB 数据库中执行查询，并将查询结果转换为字符串数组返回。
func (c *CollectionClient) QueryMany(query bson.D) ([]string, error) {
	// 使用 Find 方法在 MongoDB 中执行查询，返回一个游标和一个错误信息
	cursor, err := c.client.Find(context.TODO(), query)
	// 如果在执行查询时发生错误，返回 nil 和错误信息
	if err != nil {
		return nil, err
	}
	// 定义一个字符串数组用于存储查询结果
	var results []string
	// 使用游标的 Next 方法遍历查询结果
	for cursor.Next(context.TODO()) {
		// 将每个查询结果转换为字符串并添加到结果数组中
		results = append(results, cursor.Current.String())
	}
	// 返回查询结果和 nil 错误信息
	return results, nil
}
