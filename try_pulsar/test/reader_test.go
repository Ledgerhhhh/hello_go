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
		_ = fmt.Errorf("service error: %s\n", err)
	}
}

// 读取器(读取所有的消息,包括确认的消息)
func TestReader(t *testing.T) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: myconfig.GConfig.PulsarConfig.BrokerURL,
	})
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	reader, err := client.CreateReader(pulsar.ReaderOptions{
		Topic:          myconfig.GConfig.PulsarConfig.Topic,
		StartMessageID: pulsar.EarliestMessageID(),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	for reader.HasNext() {
		msg, err := reader.Next(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
			msg.ID(), string(msg.Payload()))
	}
}
