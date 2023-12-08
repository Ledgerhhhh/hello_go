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

func (h *MyHandler) HandleMessage(message *nsq.Message) error {
	fmt.Printf("Received message: %s\n", message.Body)
	return nil
}
func main() {
	config := nsq.NewConfig()
	config.AuthSecret = myconfig.GConfig.NsqdConfig.AuthSecret
	consumer, err := nsq.NewConsumer(myconfig.GConfig.NsqdConfig.Topic, myconfig.GConfig.NsqdConfig.Channel, config)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	handler := &MyHandler{}
	consumer.AddHandler(handler)
	err = consumer.ConnectToNSQD(myconfig.GConfig.NsqdConfig.Host +
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
