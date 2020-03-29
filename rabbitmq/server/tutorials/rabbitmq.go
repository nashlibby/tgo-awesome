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

/**
 ** 交换机定义
 ** @param durable 消息是否持久保存
 ** @param autoDelete 当最后一个消费者断开连接之后交换机是否自动被删除
 ** @param internal 设置是否内置，true为内置。如果是内置交换器，客户端无法发送消息到这个交换器中，只能通过交换器路由到交换器这种方式
 ** @param noWait 是否非阻塞，true表示是。阻塞：表示创建交换器的请求发送后，阻塞等待Server返回信息。非阻塞：不会阻塞等待Server的返回信息
 */
func (r *RabbitMq) ExchangeDeclare(exchangeName, exchangeType string, durable, autoDelete, internal, noWait bool) {
	err := r.ch.ExchangeDeclarePassive(exchangeName, exchangeType, durable, autoDelete, internal, noWait, nil)
	if err != nil {
		err = r.ch.ExchangeDeclare(
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
}

/**
 ** 队列定义
 ** @param durable 消息是否持久保存
 ** @param autoDelete 当最后一个消费者断开连接之后队列是否自动被删除
 ** @param exclusive 是否排外 1.当连接关闭时该队列是否会自动删除 2.如果不是排外的，可以使用两个消费者都访问同一个队列，如果是排外的，会对当前队列加锁，其他通道channel是不能访问的
 ** @param noWait 是否非阻塞，true表示是。阻塞：表示创建交换器的请求发送后，阻塞等待Server返回信息。非阻塞：不会阻塞等待Server的返回信息
 */
func (r *RabbitMq) QueueDeclare(queueName string, durable, autoDelete, exclusive, noWait bool) *amqp.Queue{
	queue, err := r.ch.QueueDeclarePassive(queueName, durable, autoDelete, exclusive, noWait, nil)
	if err != nil {
		queue, err = r.ch.QueueDeclare(
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
	}

	return &queue
}

/**
 ** 队列绑定
 ** @param noWait 是否非阻塞，true表示是。阻塞：表示创建交换器的请求发送后，阻塞等待Server返回信息。非阻塞：不会阻塞等待Server的返回信息
 */
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

/**
 ** 消费
 ** @param autoAck 是否自动回复ACK，true为是，回复ACK表示告诉服务器我收到消息了
 ** @param exclusive 是否排外 1.当连接关闭时该队列是否会自动删除 2.如果不是排外的，可以使用两个消费者都访问同一个队列，如果是排外的，会对当前队列加锁，其他通道channel是不能访问的
 ** @param noLocal 如果为true，表示生产者和消费者不能是同一个connect
 ** @param noWait 是否非阻塞，true表示是。阻塞：表示创建交换器的请求发送后，阻塞等待Server返回信息。非阻塞：不会阻塞等待Server的返回信息
 */
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
			// 如果autoAck为false则手动ack
			// 注意： 这里只要调用了ack就是手动确认模式，
			// multiple 表示的是在此channel中先前所有未确认的msg都将被确认
			// 并不是表示设置为false就不进行当前ack确认
			// msg.Ack(true)
			log.Printf("[x] Received a message: %s", msg.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
