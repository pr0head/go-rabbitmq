package helper

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/wagslane/go-rabbitmq"
	"strconv"
)

const (
	protobufContentType = "application/protobuf"
	jsonContentType     = "application/json"
)

func PublishProtoMessage(
	publisher rabbitmq.PublisherInterface,
	exchange string,
	routingKey []string,
	msg proto.Message,
	table rabbitmq.Table,
	expiration int,
) error {
	body, err := proto.Marshal(msg)

	if err != nil {
		return err
	}

	return publishRabbitMqMessage(
		publisher,
		exchange,
		routingKey,
		protobufContentType,
		body,
		table,
		expiration,
	)
}

func PublishJsonMessage(
	publisher rabbitmq.PublisherInterface,
	exchange string,
	routingKey []string,
	msg interface{},
	table rabbitmq.Table,
	expiration int,
) error {
	body, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	return publishRabbitMqMessage(
		publisher,
		exchange,
		routingKey,
		jsonContentType,
		body,
		table,
		expiration,
	)
}

func publishRabbitMqMessage(
	publisher rabbitmq.PublisherInterface,
	exchange string,
	routingKey []string,
	contentType string,
	body []byte,
	table rabbitmq.Table,
	expiration int,
) error {
	opts := []func(*rabbitmq.PublishOptions){
		rabbitmq.WithPublishOptionsContentType(contentType),
		rabbitmq.WithPublishOptionsPersistentDelivery,
	}

	if exchange != "" {
		opts = append(opts, rabbitmq.WithPublishOptionsExchange(exchange))
	}

	if expiration > 0 {
		opts = append(opts, rabbitmq.WithPublishOptionsExpiration(strconv.Itoa(expiration)))
	}

	if len(table) > 0 {
		opts = append(opts, rabbitmq.WithPublishOptionsHeaders(table))
	}

	return publisher.Publish(
		body,
		routingKey,
		opts...,
	)
}
