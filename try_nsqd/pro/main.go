package main

import (
	"com.ledger.goproject/myconfig"
	"errors"
	"fmt"
	"github.com/nsqio/go-nsq"
)

func init() {
	err := myconfig.InitGConfig()
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	client, err := GetProduceClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 10; i++ {
		err = client.Publish(myconfig.GConfig.NsqdConfig.Topic, []byte("hello world"))
		if err != nil {
			fmt.Println("err:", err)
			return
		}
	}
}

// GetProduceClient 获取一个生产者客户端连接
func GetProduceClient() (*nsq.Producer, error) {
	config := nsq.NewConfig()
	config.AuthSecret = myconfig.GConfig.NsqdConfig.AuthSecret

	producer, err := nsq.NewProducer(myconfig.GConfig.NsqdConfig.Host+
		":"+
		fmt.Sprintf("%d", myconfig.GConfig.NsqdConfig.Port),
		config,
	)
	if err != nil {
		return nil, errors.New("GetProduceClient error: " + err.Error())
	}
	return producer, nil
}
