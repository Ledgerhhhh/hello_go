package main

import (
	"com.ledger.goproject/myconfig"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
)

func init() {
	_ = myconfig.InitGConfig()
}

func main() {
	//TestInert()
	//TestRead()
	//TestDel()
	//TestUpdate()
	//TestInertMany()
}

func TestInertMany() {
	client := initClient()
	// 延迟关闭连接
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	collection := client.Database(myconfig.GConfig.MongoDBConfig.DataBase).
		Collection(myconfig.GConfig.MongoDBConfig.Collection)
	p1 := person{
		Name: "心",
		Age:  1,
	}
	p2 := person{
		Name: "111",
		Age:  1,
	}
	marshal, err := bson.Marshal(p1)
	if err != nil {
		return
	}
	marshal2, err := bson.Marshal(p2)
	if err != nil {
		return
	}
	i := []interface{}{marshal, marshal2}
	many, err := collection.InsertMany(context.TODO(), i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(many.InsertedIDs)
}

func TestUpdate() {
	client := initClient()
	// 延迟关闭连接
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	collection := client.Database(myconfig.GConfig.MongoDBConfig.DataBase).
		Collection(myconfig.GConfig.MongoDBConfig.Collection)
	p := person{
		Name: "ledgerhhhh",
		Age:  1,
	}
	p2 := person{
		Name: "ledgerhhh",
		Age:  1,
	}
	update := bson.D{{"$set", p2}}
	// 使用UpdateOne方法更新符合条件的第一条文档
	updateResult, err := collection.UpdateOne(context.TODO(), p, update)
	if err != nil {
		log.Fatal(err)
	}
	// 更新的文档数
	fmt.Println(updateResult.UpsertedCount)

}

func TestDel() {
	client := initClient()
	// 延迟关闭连接
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database(myconfig.GConfig.MongoDBConfig.DataBase).
		Collection(myconfig.GConfig.MongoDBConfig.Collection)
	p := person{
		Name: "ledger",
		Age:  1,
	}
	marshal, err := bson.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	one, err := collection.DeleteMany(context.TODO(), marshal)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 查询删除的文档数
	fmt.Printf("Deleted %v documents\n", one.DeletedCount)
}
func TestRead() {
	client := initClient()
	collection := client.
		Database(myconfig.GConfig.MongoDBConfig.DataBase).
		Collection(myconfig.GConfig.MongoDBConfig.Collection)
	p := person{
		Name: "ledger",
		Age:  1,
	}

	filter, err := bson.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 根据过滤条件进行模糊匹配查询
	find, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 迭代游标
	for find.Next(context.Background()) {
		var result person
		err := find.Decode(&result)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Found person: %+v\n", result)
		fmt.Println()
	}
}

type person struct {
	Name string
	Age  int
}

func TestInert() {
	client := initClient()
	collection := client.
		Database(myconfig.GConfig.MongoDBConfig.DataBase).
		Collection(myconfig.GConfig.MongoDBConfig.Collection)

	p := person{
		Name: "ledger",
		Age:  1,
	}
	res, err := collection.InsertOne(context.TODO(), p)
	if err != nil {
		return
	}
	fmt.Println(res)
}
func helloMongoDb() {
	client := initClient()
	// 选择一个数据库和一个集合
	database := client.Database(myconfig.GConfig.MongoDBConfig.DataBase)
	collection := database.Collection(myconfig.GConfig.MongoDBConfig.Collection)
	// 需要插入的数据
	document := map[string]interface{}{
		"key1":   "value1",
		"key2":   "value2",
		"key23":  "value2",
		"key233": "value2",
		"key3":   123,
	}
	// 插入一条数据
	one, err := collection.InsertOne(context.TODO(), document)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Inserted document with ID:", one.InsertedID)

	// 断开连接
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection closed.")
}

func initClient() *mongo.Client {
	// 创建一个mongodb的配置选项,配置地址,端口,用户名和密码
	clientOptions := options.Client().ApplyURI("mongodb://" +
		myconfig.GConfig.MongoDBConfig.Host +
		":" +
		strconv.Itoa(myconfig.GConfig.MongoDBConfig.Port),
	).SetAuth(options.Credential{
		// 设置用户名和密码
		Username: myconfig.GConfig.MongoDBConfig.Username,
		Password: myconfig.GConfig.MongoDBConfig.Password,
	})
	// 使用配置创建一个客户端
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("Connected to MongoDB!")
	return client
}
