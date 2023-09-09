package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://root:Qazwsx09!@res360.cn:5672//")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // 队列名称
		false,   // 是否持久化
		false,   // 是否自动删除
		false,   // 是否具有排他性
		false,   // 是否阻塞
		nil,     // 额外参数
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	err = ch.Publish(
		"",     // 交换机名称
		q.Name, // 队列名称
		false,  // 是否强制性发送
		false,  // 是否立即发送
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")

	fmt.Println("Message sent:", body)
}
