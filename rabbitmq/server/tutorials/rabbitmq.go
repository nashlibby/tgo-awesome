/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package tutorials

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMq struct {
	ch *amqp.Channel
	conn *amqp.Connection
}

// 配置信息
type ServerConfig struct {
	UserName string
	Password string
	Host     string
	Port     int64
	VHost    string
}

// 构造
func NewServer(config *ServerConfig) *RabbitMq {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/%s", config.UserName, config.Password, config.Host,
		config.Port, config.VHost))
	if err != nil {
		panic(err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err.Error())
	}

	return &RabbitMq{conn: conn, ch: ch}
}

// 关闭连接
func (r *RabbitMq) Close() {
	r.ch.Close()
	r.conn.Close()
}

// 交换机定义
func (r *RabbitMq) ExchangeDeclare(exchangeName, exchangeType string, durable, autoDelete, internal, noWait bool) {
	err := r.ch.ExchangeDeclare(
		exchangeName,   // name
		exchangeType, // type
		durable,     // durable
		autoDelete,    // auto-deleted
		internal,    // internal
		noWait,    // no-wait
		nil,      // arguments
	)

	if err != nil {
		panic(err.Error())
	}
}

// 队列定义
func (r *RabbitMq) QueueDeclare(queueName string, durable, autoDelete, exclusive, noWait bool) *amqp.Queue{
	queue, err := r.ch.QueueDeclare(
		queueName,
		durable,
		autoDelete,
		exclusive,
		noWait,
		nil,
	)
	if err != nil {
		panic(err.Error())
	}

	return &queue
}

// 队列绑定
func (r *RabbitMq) QueueBind(queueName, routingKey, exchangeName string, noWait bool) {
	err := r.ch.QueueBind(
		queueName, // queue name
		routingKey,     // routing key
		exchangeName, // exchange
		noWait,
		nil)
	if err != nil {
		panic(err.Error())
	}
}

// 消费
func (r *RabbitMq) Consume(queueName, consumerName string, autoAck, exclusive, noLocal, noWait bool) {
	msgList, err := r.ch.Consume(
		queueName, // queue
		consumerName,     // consumer
		autoAck,   // auto-ack
		exclusive,  // exclusive
		noLocal,  // no-local
		noWait,  // no-wait
		nil,    // args
	)
	if err != nil {
		panic(err.Error())
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgList {
			log.Printf("[x] Received a message: %s", msg.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
