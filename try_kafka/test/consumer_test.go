package test

import (
	"fmt"
	"github.com/IBM/sarama"
	"testing"
	"time"
)

func TestConsumer(t *testing.T) {
	testConsumer()
}
func createConsumer(brokers []string) (sarama.Consumer, error) {
	// 创建一个新的 Sarama 配置实例
	config := sarama.NewConfig()
	// 自动提交偏移量设置为每秒自动提交一次
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second
	// 分区一次拉取消息的最大字节数，控制每个分区一次能够拉取的最大字节数
	config.Consumer.Fetch.Max = 32 * 1024 // 32 KB
	// 配置通道缓冲区的大小，用于从 Kafka 服务器异步接收消息
	config.ChannelBufferSize = 1024
	// 使用配置创建 Kafka 消费者实例，返回消费者接口和可能的错误
	return sarama.NewConsumer(brokers, config)
}

func testConsumer() {
	brokers := []string{"60.204.241.30:9092"}
	consumer, err := createConsumer(brokers)
	if err != nil {
		fmt.Println(err)
	}
	partition, err := consumer.ConsumePartition("my_topic", 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Println(err)
	}
	for {
		msg := <-partition.Messages()
		// 查看分区,偏移量,key和value
		fmt.Printf("Partition: %d, Offset: %d, Key: %s, Value: %s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		time.Sleep(100 * time.Millisecond)
	}
}
