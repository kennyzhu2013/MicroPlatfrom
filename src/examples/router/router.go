package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-os/router"

	proto "github.com/micro/router-srv/proto/router"
)

var (
	routines = 3
	requests = 100
	service  = "go.micro.srv.router"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func selector(id int, r router.Router) map[string]int {
	stats := make(map[string]int)

	// select the service
	next, err := r.Select(service)
	if err != nil {
		fmt.Println(id, "error selecting", err)
		return stats
	}

	for i := 1; i <= requests; i++ {
		// get a node
		node, err := next()
		if err != nil {
			fmt.Println(id, "error getting next", err)
			return stats
		}

		stats[node.Id]++

		// make some request
		// client.Call(foo, request)
		req := client.NewRequest(service, "Router.Stats", &proto.StatsRequest{})

		var dur time.Duration

		// lets set an error
		if d := (rand.Int() % i); d == 0 {
			dur = time.Millisecond * time.Duration(rand.Int()%20)
			err = errors.InternalServerError(service, "err")
		} else if d == 1 {
			dur = time.Second * 5
			err = errors.New(service, "timed out", 408)
		} else {
			dur = time.Millisecond * time.Duration(rand.Int()%10)
			err = nil
		}

		// mark the result
		r.Mark(service, node, err)

		// record timing
		r.Record(req, node, dur, err)

		//fmt.Println(id, "selected", node.Id)
		time.Sleep(time.Millisecond*10 + time.Duration(rand.Int()%10))
	}

	return stats
}

func main() {
	cmd.Init()

	r := router.NewRouter()
	defer r.Close()

	ch := make(chan map[string]int, routines)
	stats := make(map[string]int)

	for i := 0; i < routines; i++ {
		go func(j int) {
			st := selector(j, r)
			str := ""
			for k, v := range st {
				str += fmt.Sprintf("stats %d %s %d\n", j, k, v)
			}
			fmt.Println(str)
			ch <- st
		}(i)
	}

	// collate stats
	for i := 0; i < routines; i++ {
		st := <-ch
		for k, v := range st {
			stats[k] += v
		}
	}

	for k, v := range stats {
		fmt.Println("overall stats", k, v)
	}
}
