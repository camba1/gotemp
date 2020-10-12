package globalUtils

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/micro/go-micro/v2/broker"
	"log"
)

//MyBroker is a Struct type that contains all the broker functionality
type MyBroker struct {
	Br broker.Broker
}

//ProtoToByte converts a proto message to a byte slice so that it can be sent out to the broker
func (mb *MyBroker) ProtoToByte(protoMsg proto.Message) ([]byte, error) {
	byteUser, err := proto.Marshal(protoMsg)
	if err != nil {
		log.Printf(glErr.BrkBadMarshall(fmt.Sprintf("User: %v", protoMsg), err))
		return nil, err
	}
	return byteUser, nil
}

//SendMsg sends message to broker so that is can be picked up by a subscription at some point. This is setup to be fire and forget
func (mb *MyBroker) SendMsg(objectToSend []byte, header map[string]string, topic string) error {

	var message broker.Message
	message.Header = header
	message.Body = objectToSend

	if mb.Br == nil {
		err1 := errors.New("broker is not defined or connected")
		log.Printf(glErr.BrkNoMessageSent(topic, err1))
		return err1

	}
	err := mb.Br.Connect()
	if err != nil {
		//log.Printf("unable to connect to broker: Error: %v" , err)
		return err
	}

	err = mb.Br.Publish(topic, &message)
	if err != nil {
		log.Printf(glErr.BrkNoMessageSent(topic, err))
		return err
	}
	log.Printf("sent message to Topic %s. Message %v", topic, &header)
	return nil
}

//SubToMsg subscribes to a message in the broker. to pick up all messages, leave the queueName empty, otherwise
// message will be sent to just one of subscribers with that queueName
func (mb *MyBroker) SubToMsg(subHandler broker.Handler, topic string, queueName string) error {

	err := mb.Br.Connect()
	if err != nil {
		log.Printf(glErr.BrkNoConnection(err))
		return err
	}

	if queueName != "" {
		_, err = mb.Br.Subscribe(topic, subHandler, broker.Queue(queueName))
	} else {
		_, err = mb.Br.Subscribe(topic, subHandler)
	}
	if err != nil {
		log.Printf(glErr.BrkUnableToSetSubs(topic, err))
		return err
	}
	log.Printf("Subscribed to queue: %s, queue: %s", topic, queueName)
	return nil
}

//GetMsg breaks the message from the broker into its three component parts
func (mb *MyBroker) GetMsg(p broker.Event) (string, map[string]string, []byte, error) {
	fmt.Printf("Received message for subscription Topic %s: %v\n", p.Topic(), p.Message().Header)
	return p.Topic(), p.Message().Header, p.Message().Body, nil
}
