package main

import (
	"com.ledger.goproject/myconfig"
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	err := myconfig.InitGConfig()
	if err != nil {
		fmt.Println("err", err)
	}
}

type MyHandler struct{}

func (m MyHandler) HandleMessage(message *nsq.Message) error {
	fmt.Printf("Received message: %s\n", message.Body)
	return nil
}

func main() {
	// 创建一个消费者,定义主题和通道
	consumer, err := nsq.NewConsumer(
		myconfig.GConfig.NsqdConfig.Topic,
		myconfig.GConfig.NsqdConfig.Channel,
		nsq.NewConfig(),
	)

	if err != nil {
		fmt.Println("err", err)
		return
	}
	handler := &MyHandler{}
	consumer.AddHandler(handler)
	err = consumer.ConnectToNSQD(
		myconfig.GConfig.NsqdConfig.Host +
			":" +
			fmt.Sprintf("%d", myconfig.GConfig.NsqdConfig.Port),
	)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	// 等待程序退出信号
	<-shutdown

	// 停止消费者
	consumer.Stop()
	fmt.Println("Shutting down...")
}
