/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package main

import (
	"log"
	"rabbitmq/tutorials"
)

func main() {
	config := tutorials.ClientConfig{
		UserName: "guest",
		Password: "guest",
		Host:     "localhost",
		Port:     5672,
		VHost:    "/",
	}

	client := tutorials.NewClient(&config)
	defer client.Close()

	// simple模式 直接队列投递 多消费者轮询消费
	queueName := "simple.queue"
	body := "simple message"
	queue := client.QueueDeclare(queueName, false, false, false, false)
	err := client.Publish([]byte(body), "", queue.Name, false, false)
	if err != nil {
		log.Printf(" [x] Send %s failed", body)
	} else {
		log.Printf(" [x] Send %s succeeded", body)
	}

	// // fanout模式 类似广播 转发到所有绑定交换机的queue
	// exchangeName := "fanout.exchange"
	// exchangeType := "fanout"
	// queueName := ""
	// body := "fanout message"
	//
	// client.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false)
	// queue := client.QueueDeclare(queueName, false, false, false, false)
	// err := client.Publish([]byte(body), exchangeName, queue.Name, false, false)
	// if err != nil {
	// 	log.Printf(" [x] Send %s failed", body)
	// } else {
	// 	log.Printf(" [x] Send %s succeeded", body)
	// }

	// // direct模式 类似单播 routingKey和bindingKey完全匹配
	// exchangeName := "direct.exchange"
	// exchangeType := "direct"
	// routingKey := "direct.routing.key"
	// body := "direct message"
	//
	// client.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false)
	// err := client.Publish([]byte(body), exchangeName, routingKey, false, false)
	// if err != nil {
	// 	log.Printf(" [x] Send %s failed", body)
	// } else {
	// 	log.Printf(" [x] Send %s succeeded", body)
	// }

	// // topic模式 类型组播 转发到符合通配符匹配的queue
	// exchangeName := "topic.exchange"
	// exchangeType := "topic"
	// routingKey := "topic.routing.key.demo"
	// body := "topic message"
	//
	// client.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false)
	// err := client.Publish([]byte(body), exchangeName, routingKey, false ,false)
	// if err != nil {
	// 	log.Printf(" [x] Send %s failed", body)
	// } else {
	// 	log.Printf(" [x] Send %s succeeded", body)
	// }
}
