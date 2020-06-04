package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/micro/go-micro/v2/broker"
	pb "goTemp/user/proto"
	"log"
)

type myBroker struct {
	br broker.Broker
}

var mb myBroker

func (mb *myBroker) sendMsg(user *pb.User, header map[string]string, topic string) error {
	var message broker.Message
	binUser, err := proto.Marshal(user)
	if err != nil {
		log.Printf(glErr.BrkBadMarshall(fmt.Sprintf("User: %v", user), err))
		return err
	}
	message.Header = header
	message.Body = binUser

	//err = br.Connect()
	//if err != nil {
	//	log.Printf("unable to connect: Error: %v" , err)
	//	return err
	//}

	err = mb.br.Publish(topic, &message)
	if err != nil {
		log.Printf(glErr.BrkNoMessageSent(topic, err))
		return err
	}
	log.Printf("sent message to topic %s. Message %v", topic, &user)
	return nil
}

func (mb *myBroker) subToMsg(subHandler broker.Handler, topic string, queueName string) error {

	err := mb.br.Connect()
	if err != nil {
		log.Printf(glErr.BrkNoConnection(err))
		return err
	}

	//var subs broker.Subscriber

	if queueName != "" {
		_, err = mb.br.Subscribe(topic, subHandler, broker.Queue(queueName))
	} else {
		_, err = mb.br.Subscribe(topic, subHandler)
	}
	if err != nil {
		log.Printf(glErr.BrkUnableToSetSubs(topic, err))
		return err
	}
	log.Printf("Subscribed to queue: %s, queue: %s", topic, queueName)
	return nil
}

func getMsg(p broker.Event) error {
	receivedUser := pb.User{}
	err := proto.Unmarshal(p.Message().Body, &receivedUser)
	if err != nil {
		log.Printf(glErr.BrkBadUnMarshall(p.Topic(), p.Message().Body, err))
		return err
	}
	fmt.Printf("Received user for subscription topic %s: %v\n", p.Topic(), &receivedUser)
	return nil
}
