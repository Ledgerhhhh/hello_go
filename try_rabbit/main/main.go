package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"sync"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:admin@106.54.9.19:5672/ledger")
	if err != nil {
		fmt.Println(err)
		return
	}
	//failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 创建一个通道
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ch.Close()
	// 定义一个交换机
	err = ch.ExchangeDeclare(
		"test",   // 交换机名称
		"fanout", // 交换机类型
		true,     // 持久性
		false,    // 自动删除
		false,    // 内部使用
		false,    // 不等待服务器响应
		nil,      // 参数
	)
	// 定义一个队列
	q, err := ch.QueueDeclare(
		"my_queue", // 队列名称
		false,      // 持久性
		false,      // 自动删除
		false,      // 独占性
		false,      // 不等待服务器响应
		nil,        // 参数
	)
	// 交换机和队列的绑定
	err = ch.QueueBind(
		q.Name, // 队列名称
		"",     // 路由键（在fanout交换机类型中通常为空）
		"test", // 交换机名称
		false,  // 不等待服务器响应
		nil,    // 参数
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 发送消息到队列
	body := "Hello, RabbitMQ!"
	group := sync.WaitGroup{}
	//k := 10
	//k := 50
	k := 100
	group.Add(k)
	for i := 0; i < k; i++ {
		go func() {
			defer group.Done()
			err = ch.Publish(
				"test", // 交换机名称
				"",     // 路由键
				false,  // 强制
				false,  // 立即发送
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				},
			)
			fmt.Printf(" [x] Sent %s\n", body)
		}()
	}
	group.Wait()
	if err != nil {
		fmt.Println(err)
		return
	}

}
