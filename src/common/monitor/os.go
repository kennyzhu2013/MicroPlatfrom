/*
@Time : 2019/4/23 15:42 
@Author : kenny zhu
@File : os.go
@Software: GoLand
@Others:
*/
package monitor

import (
	"context"
	// "github.com/micro/go-micro/registry"
	"registry"

	"log/log"
	"strconv"
	"sync"
	"time"
)

type os struct {
	exit chan bool

	registry registry.Registry
	opts    *Options
	next    func() []string

	sync.RWMutex
	// bHeart  bool
	// cache.
	heartbeats map[string]Heartbeat
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
		dOpts.Interval = HeartBeatCheck
	}

	o := &os{
		registry:    newRegistry(opts...),
		opts:       dOpts,
		exit:       make(chan bool),
		heartbeats: make(map[string]Heartbeat, 1000),
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
	defer func() {
		if v := recover();v != nil {
			log.Errorf("heartbeat panic:%v", v)
		}
	}()

	for {
		select {
		case <-t.C:
			var heartbeats [] Heartbeat
			timeNow := time.Now().Unix()
			o.Lock()
			for _, hb := range o.heartbeats {
				// get not out-dated heartbeats
				if hb.timestamp + hb.ttl < timeNow {
					heartbeats = append(heartbeats, hb)
				}
			}

			// clear heartbeats cache.
			if len(o.heartbeats) > 1000 {
				o.heartbeats = make(map[string]Heartbeat, 1000)
			}
			o.Unlock()
			log.Infof("Heartbeats length:%v", len(heartbeats))
			for _, hb := range heartbeats {
				hb.timestamp = timeNow

				// get node.
				for _, node:= range hb.service.Nodes {
					if node.Id == hb.id {
						node.Metadata["timestamp"] = strconv.FormatInt(hb.timestamp,10)
						status := node.Metadata[ServiceStatus]
						if hb.weights < DownLimits {
							status = NormalState
						} else if hb.weights > UpLimits {
							status = UpThresh
						} else {
							// nothing to do...
						}
						node.Metadata[ServiceStatus] = status

						// pub := o.opts.Client.NewPublication(HeartbeatTopic, hb)
						err := registry.Register( hb.service, registry.RegisterTTL(o.opts.Interval)  )
						if err != nil {
							log.Fatal("Heartbeats check failed!")
							//  fmt.Printf("Heartbeats check failed!\n")
						}
						log.Infof("Heartbeats check success,node:%v", node.Id)
						//  fmt.Printf("Heartbeats check success,node:%v\n", node.Id)
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
	hb := Heartbeat{
		id:       service.Nodes[0].Id,
		service:  service,
		ttl:      int64((o.opts.Interval.Seconds()) * 2),
		// timestamp: time.Now().Unix(),
	}

	o.heartbeats[hb.id] = hb
	o.Unlock()
}

// push heartbeat to service directly..
func UpdateServiceWithoutSync(service *registry.Service, weights int)  {
	timeNow := time.Now().Unix()

	// service nodes register, here only one
	for _, node:= range service.Nodes {
		node.Metadata["timestamp"] = strconv.FormatInt(timeNow, 10)
		status := node.Metadata[ServiceStatus]
		if weights < DownLimits {
			status = NormalState
		} else if weights > UpLimits {
			status = UpThresh
		} else {
			// nothing to do...
		}
		node.Metadata[ServiceStatus] = status
	}

	// else to do?..
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
