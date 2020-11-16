package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-os/event"

	"golang.org/x/net/context"
)

var (
	types = []string{"Provisioned", "Restarted", "Deprovisioned"}
)

type Data struct {
	Foo string
	Bar int64
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randType() string {
	return types[rand.Int()%len(types)]
}

func pub(ev event.Event) {
	origin := "ev-client"
	rootId := fmt.Sprintf("root.%s.%d", origin, time.Now().UnixNano())

	b, _ := json.Marshal(&Data{
		Foo: "foo",
		Bar: time.Now().Unix(),
	})

	for i := 0; i < 10; i++ {
		r := &event.Record{
			Id:        fmt.Sprintf("%s.%d", origin, time.Now().UnixNano()),
			Type:      "agent." + randType(),
			Origin:    origin,
			Timestamp: time.Now().Unix(),
			RootId:    rootId,
			Metadata: map[string]string{
				"state":   "update",
				"action":  "dosomething",
				"message": "some event",
			},
			Data: string(b),
		}

		fmt.Println("Sending event", r.Id, r.Type)

		if err := ev.Publish(context.TODO(), r); err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Millisecond * 100)
	}
}

func sub(ev event.Event) {
	// subscribe to event type
	err := ev.Subscribe(context.TODO(), func(rec *event.Record) {
		fmt.Println("Received event", rec.Id, rec.Type)
	}, "agent."+randType())

	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	cmd.Init()

	// create event handler
	ev := event.NewEvent()

	// subscribe to events
	go sub(ev)

	// publish events
	pub(ev)
}
