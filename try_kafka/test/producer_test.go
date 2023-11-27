package test

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestProducer(t *testing.T) {
	testProducer()
}

type CustomPartitioner struct{}

// 确定你的消息发送给哪个分区
func (c CustomPartitioner) Partition(message *sarama.ProducerMessage, numPartitions int32) (int32, error) {
	// 检查消息的 key 是否以 "kk" 开头
	if key, ok := message.Key.(sarama.StringEncoder); ok {
		if strings.HasPrefix(string(key), "kk") {
			// 如果以 "kk" 开头，将消息发送到分区 0
			return 0, nil
		}
	}
	return 1, nil
}

func (c CustomPartitioner) RequiresConsistency() bool {
	return false
}

func NewCustomPartitioner(topic string) sarama.Partitioner {
	return &CustomPartitioner{}
}

func createProducer(brokers []string) (sarama.AsyncProducer, error) {
	// 创建一个新的 Kafka 生产者配置
	config := sarama.NewConfig()
	// 设置消息确认级别
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 开启幂等性
	config.Producer.Idempotent = true
	// 设置生产者在消息发送成功时是否返回成功的通知
	config.Producer.Return.Successes = true
	// 设置生产者的超时时间，即在发送消息时等待成功的最大时间
	config.Producer.Timeout = 5 * time.Second
	// 指定分区器
	config.Producer.Partitioner = NewCustomPartitioner
	// 设置消息刷新的频率
	config.Producer.Flush.Frequency = 100 * time.Millisecond
	// 设置等待的未刷新消息的数量
	config.Producer.Flush.Messages = 1000
	// 设置单次刷新中最大的消息数
	config.Producer.Flush.MaxMessages = 10000
	// 设置消息压缩类型
	config.Producer.Compression = sarama.CompressionSnappy
	// 使用配置和代理地址列表创建一个异步生产者实例
	producer, err := sarama.NewAsyncProducer(brokers, config)
	// 返回创建的生产者实例和可能的错误
	return producer, err
}

func testProducer() {
	brokers := []string{"60.204.241.30:9092"}
	producer, err := createProducer(brokers)
	if err != nil {
		log.Fatal(err)
	}

	// 使用 WaitGroup 来等待所有异步操作完成
	var wg sync.WaitGroup
	wg.Add(1)
	for i := 0; i < 52; i++ {
		msg := "hello kafka" + strconv.Itoa(i)
		producer.Input() <- &sarama.ProducerMessage{
			Topic:     "my_topic",
			Partition: 0,
			Key:       sarama.StringEncoder(fmt.Sprintf("%s %d", msg, i)),
			Value:     sarama.StringEncoder(fmt.Sprintf("%s %d", msg, i)),
		}
	}
	// 处理成功和失败的回调
	go func() {
		for {
			select {
			case success := <-producer.Successes():
				fmt.Printf("Message sent to partition %d at offset %d\n", success.Partition, success.Offset)
			case err := <-producer.Errors():
				fmt.Printf("Failed to send message: %v\n", err.Err)
			}
		}
	}()
	// 等待所有异步操作完成
	wg.Wait()
	// 关闭生产者
	producer.AsyncClose()
}
