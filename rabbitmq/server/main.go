/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package main

import "rabbitmq/tutorials"

func main() {
	config := tutorials.ServerConfig{
		UserName: "guest",
		Password: "guest",
		Host:     "localhost",
		Port:     5672,
		VHost:    "/",
	}

	server := tutorials.NewServer(&config)
	defer server.Close()

	// // simple模式 直接队列投递 多消费者轮询消费
	// queueName := "simple-queue"
	// queue := server.QueueDeclare(queueName, false, false, false, false)
	// server.Consume(queue.Name, "", true, false, false, false)

	// direct模式 类似单播 routingKey和bindingKey完全匹配
	exchangeName := "direct-exchange"
	exchangeType := "direct"
	routingKey := "direct-routing-key"

	server.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false)
	queue := server.QueueDeclare("", false, false, true, false)
	server.QueueBind(queue.Name, routingKey, exchangeName, false)
	server.Consume(queue.Name, "", true, false, false, false)

	// fanout模式 类似广播 转发到所有绑定交换机的queue
	// exchangeName := "fanout-exchange"
	// exchangeType := "fanout"
	// routingKey := ""
	// queueName := ""
	//
	// server.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false)
	// queue := server.QueueDeclare(queueName, false, false, true, false)
	// server.QueueBind(queue.Name, routingKey, exchangeName, false)
	// server.Consume(queue.Name, "", true, false, false, false)
}
