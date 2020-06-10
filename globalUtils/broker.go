package globalUtils

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/micro/go-micro/v2/broker"
	"log"
)

type MyBroker struct {
	Br broker.Broker
}

//var mb MyBroker

func (mb *MyBroker) ProtoToByte(protoMsg proto.Message) ([]byte, error) {
	byteUser, err := proto.Marshal(protoMsg)
	if err != nil {
		log.Printf(glErr.BrkBadMarshall(fmt.Sprintf("User: %v", protoMsg), err))
		return nil, err
	}
	return byteUser, nil
}

func (mb *MyBroker) SendMsg(objectToSend []byte, header map[string]string, topic string) error {

	var message broker.Message
	message.Header = header
	message.Body = objectToSend

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

func (mb *MyBroker) SubToMsg(subHandler broker.Handler, topic string, queueName string) error {

	err := mb.Br.Connect()
	if err != nil {
		log.Printf(glErr.BrkNoConnection(err))
		return err
	}

	//var subs broker.Subscriber

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

func (mb *MyBroker) GetMsg(p broker.Event) (string, map[string]string, []byte, error) {
	//var receivedMsg proto.Message
	//log.Printf("unmarshalling message: %v", p.Message().Header)
	//err := proto.Unmarshal(p.Message().Body, receivedMsg)
	//if err != nil {
	//	log.Printf(glErr.BrkBadUnMarshall(p.Topic(), p.Message().Body, err))
	//	return "", nil, nil, err
	//}
	fmt.Printf("Received message for subscription Topic %s: %v\n", p.Topic(), p.Message().Header)
	//return p.Topic(), p.Message().Header, &receivedMsg, nil
	return p.Topic(), p.Message().Header, p.Message().Body, nil
}
