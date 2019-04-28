/*
@Time : 2019/4/28 14:37 
@Author : kenny zhu
@File : test_broker.go
@Software: GoLand
@Others:
*/
package test

import (
	"time"
	"fmt"
	"log"

	"rabbitmq"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/broker"
)

var (
	topic = "go.micro.topic.foo"
	rbroker = rabbitmq.NewBroker()
)

func pub() {
	tick := time.NewTicker(time.Second)
	i := 0
	for _ = range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}
		if err := rbroker.Publish(topic, msg); err != nil {
			log.Printf("[pub] failed: %v", err)
		} else {
			fmt.Println("[pub] pubbed message:", string(msg.Body))
		}
		i++
	}
}

func sub() {
	_, err := rbroker.Subscribe(topic, func(p broker.Publication) error {
		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	cmd.Init()

	if err := rbroker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := rbroker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	go pub()
	go sub()

	<-time.After(time.Second * 10)
}
