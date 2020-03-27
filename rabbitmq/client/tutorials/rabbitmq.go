/**
 * Created by nash.tang.
 * Author: nash.tang <112614251@qq.com>
 */

package tutorials

import (
	"fmt"
	"github.com/streadway/amqp"
)

type RabbitMq struct {
	ch *amqp.Channel
	conn *amqp.Connection
}

// 配置信息
type ClientConfig struct {
	UserName string
	Password string
	Host     string
	Port     int64
	VHost    string
}

// 构造
func NewClient(config *ClientConfig) *RabbitMq {
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

// 队列
func (r *RabbitMq) Queue(name string, durable, autoDelete, exclusive, noWait bool) *amqp.Queue{
	queue, err := r.ch.QueueDeclare(
		name,
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

// 发布
func (r *RabbitMq) Publish(body []byte, exchange, routingKey string, mandatory, immediate bool) error {
	err := r.ch.Publish(
		exchange,
		routingKey,
		mandatory,
		immediate,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	return err
}
