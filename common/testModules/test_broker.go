/*
@Time : 2019/4/28 14:37 
@Author : kenny zhu
@File : test_broker.go
@Software: GoLand
@Others:
*/
package main

import (
	"time"

	"fmt"
	"log"

	"rabbitmq"

	"broker"
	"encoding/json"
)

var (
	storeTopic = "recording.store"
	audioTopic = "recording.audio"
	rbroker = rabbitmq.NewBroker()
)

func pub() {
	tick := time.NewTicker(time.Second)
	i := 0
	for _ = range tick.C {
		var params struct {
			LeftFilename  string
			RightFilename string
		}

		params.LeftFilename = "/tmp/record/z9hG4bK10629d31d11b4syd6sz16s9w9042buu6z@139.120.40.22_early_left.amr"
		params.RightFilename = "/tmp/record/z9hG4bK10629d31d11b4syd6sz16s9w9042buu6z@139.120.40.22_early_right.amr"
		prefersJson,_ := json.Marshal( params )
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(prefersJson),
		}
		if err := rbroker.Publish(storeTopic, msg); err != nil {
			log.Printf("[pub] failed: %v", err)
		} else {
			fmt.Println("[pub] pubbed message:", string(msg.Body))
		}
		i++
	}
}

func pub2() {
	tick := time.NewTicker(time.Second)
	i := 0
	for _ = range tick.C {
		var params struct {
			CallId   string
		}

		params.CallId = "z9hG4bK10629d31d11b4syd6sz16s9w9042buu6z"
		prefersJson,_ := json.Marshal( params )
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(prefersJson),
		}
		if err := rbroker.Publish(audioTopic, msg); err != nil {
			log.Printf("[pub] failed: %v", err)
		} else {
			fmt.Println("[pub] pubbed message:", string(msg.Body))
		}
		i++
	}
}
func sub() {
	_, err := rbroker.Subscribe(storeTopic, func(p broker.Publication) error {
		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// cmd.Init()
	// broker.DefaultBroker = rbroker
	if err := rbroker.Init( broker.Addrs("amqp://admin:cmcc888@10.153.90.11:5672") ); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := rbroker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	go pub()
	go sub()
	// go pub2()

	<-time.After(time.Second * 60)
}
