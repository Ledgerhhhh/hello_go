package main

import (
	"com.ledger.goproject/myconfig"
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
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
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	return nil
}
func main() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("XDControl", "Duwi.ThirdPlatform.SmartVoice", config)
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

	for {
		time.Sleep(3 * time.Second)
		fmt.Println("sleep")
	}
}
