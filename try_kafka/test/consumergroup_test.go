package test

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"testing"
	"time"
)

func getConsumerGroup(brokers []string) (sarama.ConsumerGroup, error) {
	// 创建一个新的 Sarama 配置实例
	config := sarama.NewConfig()
	// 自动提交偏移量设置为每秒自动提交一次
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second
	// 设置从最早的偏移量开始消费
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	// 分区一次拉取消息的最大字节数，控制每个分区一次能够拉取的最大字节数
	config.Consumer.Fetch.Max = 32 * 1024 // 32 KB
	// 配置通道缓冲区的大小，用于从 Kafka 服务器异步接收消息
	config.ChannelBufferSize = 1024
	return sarama.NewConsumerGroup(brokers, "test1", config)
}

type consumerGroupHandler struct {
}

func (n consumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (n consumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (n consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		// 处理接收到的消息
		fmt.Printf("Partition: %d, Offset: %d, Key: %s, Value: %s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))

		// 标记消息已处理
		session.MarkMessage(msg, "")
	}
	return nil
}

func TestConsumerGroup(t *testing.T) {
	brokers := []string{
		"60.204.241.30:9092",
	}
	group, err := getConsumerGroup(brokers)
	if err != nil {
		fmt.Println(err)
	}
	var handler consumerGroupHandler
	topics := []string{
		"my_topic",
	}
	ctx, cancelFunc := context.WithCancel(context.Background())

	for i := 0; i < 3; i++ {
		err = group.Consume(ctx, topics, handler)
		if err != nil {
			fmt.Println(err)
		}
	}

	time.Sleep(2 * time.Second)
	group.Close()
	cancelFunc()
}
