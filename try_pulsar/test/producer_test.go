package test

import (
	"com.ledger.goproject/myconfig"
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"sync"
	"testing"
)

func init() {
	err := myconfig.InitGConfig()
	if err != nil {
		_ = fmt.Errorf("ttl_service error: %s\n", err)
	}
}

// 生产者(用于发送消息到队列)
func TestProducer(t *testing.T) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:            myconfig.GConfig.PulsarConfig.BrokerURL,
		Authentication: pulsar.NewAuthenticationToken("eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJsZWRnZXIifQ.F3pKYZPn6hsuU8sa0x-e7kBPKcpu--IJtGyzsz37OCQ"),
	})

	if err != nil {
		t.Fatalf("Failed to create Pulsar client: %v", err)
	}
	defer client.Close()

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: myconfig.GConfig.PulsarConfig.Topic,
	})
	if err != nil {
		t.Fatalf("Failed to create Pulsar producer: %v", err)
	}
	defer producer.Close()
	//k := 10
	//k := 50
	//k := 100
	k := 200
	group := sync.WaitGroup{}
	group.Add(k)

	for i := 0; i < k; i++ {
		go func() {
			_, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
				Payload: []byte("11111111"),
			})
			if err != nil {
				fmt.Printf("Failed to send message: %v\n", err)
				return
			}
			fmt.Println("Sent a message")
			group.Done()
		}()
	}

	group.Wait()

	fmt.Println("All messages sent")

	// Add additional code if needed
}

func TestT(t *testing.T) {
	group := sync.WaitGroup{}

	group.Add(10)

	for i := 0; i < 9; i++ {
		go func() {
			fmt.Println("hhh")
			group.Done()
		}()
	}

	group.Wait()
}
