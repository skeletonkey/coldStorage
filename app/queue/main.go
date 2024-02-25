package queue

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/skeletonkey/lib-core-go/logger"
)

var connection *amqp.Connection
var channel *amqp.Channel

func GetChannel() {
	config := getConfig()
	logger := logger.Get()

	var err error

	connection, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", config.User, config.Password, config.Host, config.Port))
	if err != nil {
		logger.Error().Err(err).Msg("Failed to connect to RabbitMQ")
	}

	channel, err = connection.Channel()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to open a channel")
	}

}

func PublishMsg(queueName string, msg string) error {
	logger := logger.Get()
	config := getConfig()
	logger.Trace().Str("queue", queueName).Str("msg", msg).Msg("PublishMsg")
	queueConfig, ok := config.Queues[queueName]
	if !ok {
		return fmt.Errorf("unable to find queue named %s", queueName)
	}
	q, err := channel.QueueDeclare(
		queueName,                    // name
		queueConfig.Durable,          // durable
		queueConfig.DeleteWhenUnused, // delete when unused
		queueConfig.Exclusive,        // exclusive
		queueConfig.NoWait,           // no-wait
		queueConfig.Arguments,        // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = channel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if err != nil {
		return err
	}
	logger.Trace().Msg("Message sent")

	return nil
}

func CloseConnection() {
	connection.Close()
}

func CloseChannel() {
	channel.Close()
}
