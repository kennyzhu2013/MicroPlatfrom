package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-os/monitor"
)

type request struct {
	service string
	method  string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (r *request) Service() string {
	return r.service
}

func (r *request) Method() string {
	return r.method
}

func record(m monitor.Monitor) {
	reqs := []*request{
		&request{"foo", "Foo.Bar"},
		&request{"bar", "Baz.Bar"},
		&request{"baz", "Bar.Man"},
	}

	for i := 0; i < 100; i++ {
		for _, r := range reqs {
			t := time.Duration(rand.Float64() * 1e9)

			if i%2 == 0 {
				m.RecordStat(r, t, nil)
			} else {
				m.RecordStat(r, t, errors.New("an error"))
			}
		}
		time.Sleep(time.Millisecond * 10)
	}
}

// Return successful healthcheck
func success() (map[string]string, error) {
	return map[string]string{
		"msg":    "a successful check",
		"foo":    "bar",
		"metric": "1",
		"label":  "cruft",
		"stats":  "123.0",
	}, nil
}

// Return failing healthcheck
func failure() (map[string]string, error) {
	return map[string]string{
		"msg":    "a catastrophic failure occurred",
		"foo":    "ugh",
		"metric": "-0.0001",
		"label":  "",
		"stats":  "NaN",
	}, errors.New("Unknown exception")
}

func main() {
	cmd.Init()
	m := monitor.NewMonitor(
		monitor.Interval(time.Second),
	)

	defer m.Close()

	hc1 := m.NewHealthChecker("go.micro.healthcheck.ping", "This is a ping healthcheck that succeeds", success)
	hc2 := m.NewHealthChecker("go.micro.healthcheck.pong", "This is a pong healthcheck that fails", failure)

	m.Register(hc1)
	m.Register(hc2)

	go record(m)

	<-time.After(time.Second * 10)
	fmt.Println("Stopping monitor")
}
