/*
@Time : 2019/4/23 15:42 
@Author : kenny zhu
@File : os.go
@Software: GoLand
@Others:
*/
package monitor

import (
	"github.com/micro/go-micro/registry"
	"sync"
	"context"
	"time"

	"github.com/kennyzhu/go-os/log"
	"strconv"
)

type os struct {
	exit chan bool

	registry registry.Registry
	opts    *Options
	next    func() []string

	sync.RWMutex
	// bHeart  bool
	// cache.
	heartbeats map[string] *Heartbeat
}

// use consul default.
func newOS(newRegistry NewRegistry, opts ...registry.Option) Monitor {
	if newRegistry == nil {
		// use consul, here must use et-cd instead.
		newRegistry = registry.NewRegistry
	}

	options := registry.Options{
		Context: context.Background(),
	}

	for _, o := range opts {
		o(&options)
	}
	dOpts := getOptions(options.Context)


	// set default interval, the same with registry config.
	if dOpts.Interval == time.Duration(0) {
		dOpts.Interval = time.Second * 30
	}

	o := &os{
		registry:    newRegistry(opts...),
		opts:       dOpts,
		exit:       make(chan bool),
		heartbeats: make(map[string]*Heartbeat),
		// cache:      make(map[string][]*registry.Service), // local cache?..
	}

	// default run.
	go o.run()
	return o
}

// watch to monitor
func (o *os) run() {
	// not need to watch as service.
	o.heartbeat()
}

// Send heartbeats as client every o.opts.Interval time for every register service.
// eg: service := (
// 	Name("com.example.srv.foo"),
// 	WithTTL(time.Second*30),
// 	WithInterval(time.Second*15),
// )
func (o *os) heartbeat() {
	t := time.NewTicker(o.opts.Interval)

	for {
		select {
		case <-t.C:
			var heartbeats [] *Heartbeat
			timeNow := time.Now().Unix()
			o.RLock()
			for _, hb := range o.heartbeats {
				// get not out-dated heartbeats
				if hb.timestamp + hb.ttl < timeNow {
					heartbeats = append(heartbeats, hb)
				}
			}
			o.RUnlock()

			for _, hb := range heartbeats {
				hb.timestamp = timeNow

				// get node.
				for _, node:= range hb.service.Nodes {
					if node.Id == hb.id {
						node.Metadata["timestamp"] = strconv.FormatInt(hb.timestamp,10)
						status := node.Metadata[ServiceStatus]
						if hb.weights < downLimits {
							status = NormalState
						} else if hb.weights > upLimits  {
							status = UpThresh
						} else {
							// nothing to do...
						}
						node.Metadata[ServiceStatus] = status

						// pub := o.opts.Client.NewPublication(HeartbeatTopic, hb)
						err := registry.Register( hb.service )
						if err != nil {
							log.Fatal("Heartbeats check failed!")
						}
						break;
					}
				}
			}
		case <-o.exit:
			return
		}
	}
}

// check heart beat.
func (o *os) PushHeartBeat(service *registry.Service, weights int)  {
	o.Lock()
	hb := &Heartbeat{
		id:       service.Nodes[0].Id,
		service:  service,
		ttl:      int64((o.opts.Interval.Seconds()) * 2),
	}

	o.heartbeats[hb.id] = hb
	o.Unlock()
}

func (o *os) GetMonitorRegistry() registry.Registry {
	return  o.registry
}

func (o *os) Close() error {
	select {
	case <-o.exit:
		return nil
	default:
		close(o.exit)
	}
	return nil
}
