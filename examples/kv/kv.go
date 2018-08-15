package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-os/kv"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func exit() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-ch
	os.Exit(1)
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService()
	service.Init()

	// Start the key value gossiping
	keyval := kv.NewKV()
	defer keyval.Close()

	// Start server since we have to be
	// part of the hash ring for now
	go service.Run()

	go exit()

	id := rand.Int() % 100

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("foo%d-%d", i, id)

		err := keyval.Put(&kv.Item{
			Key:        key,
			Value:      []byte(`hello`),
			Expiration: time.Second * 20,
		})
		if err != nil {
			fmt.Println("put err", err)
		}

		item, err := keyval.Get(key)
		if err != nil {
			fmt.Println("get err", err)
		} else {
			fmt.Println("get", item)
		}

		time.Sleep(time.Second)
	}
}
