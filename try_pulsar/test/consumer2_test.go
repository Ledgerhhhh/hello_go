package test

import (
	"com.ledger.goproject/myconfig"
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
	"testing"
)

func init() {
	err := myconfig.InitGConfig()
	if err != nil {
		_ = fmt.Errorf("ttl_service error: %s\n", err)
	}
}

// 消费者(读取消息,并选择是否要确认消息)
func TestConsumer2(t *testing.T) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: myconfig.GConfig.PulsarConfig.BrokerURL,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            myconfig.GConfig.PulsarConfig.Topic,
		SubscriptionName: myconfig.GConfig.PulsarConfig.SubscriptionName,
		// 订阅者的类型
		Type: pulsar.Shared,
	})

	defer consumer.Close()
	// 获取消息
	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		payload := msg.Payload()

		fmt.Printf("Received message msgId: %s -- content: '%s'\n", msg.ID(), string(payload))
		err = consumer.Ack(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}
