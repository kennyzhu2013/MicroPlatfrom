package router

import (
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/selector"
	"github.com/micro/go-micro/server"

	proto "github.com/micro/router-srv/proto/router"

	"golang.org/x/net/context"
)

type os struct {
	exit chan bool
	opts selector.Options

	client client.Client
	server server.Server

	r proto.RouterService

	// TODO
	// selector cache service:[]versions
	// stats cache client calls service:[]stats

	sync.RWMutex
	stats map[string]*stats
	cache map[string]*cache
}

var (
	publishInterval = time.Second * 10
)

func newOS(opts ...selector.Option) Router {
	options := selector.Options{
		Context: context.TODO(),
	}

	for _, o := range opts {
		o(&options)
	}

	c, ok := client.FromContext(options.Context)
	if !ok {
		c = client.DefaultClient
	}

	s, ok := server.FromContext(options.Context)
	if !ok {
		s = server.DefaultServer
	}

	o := &os{
		exit:   make(chan bool),
		opts:   options,
		client: c,
		server: s,
		cache:  make(map[string]*cache),
		stats:  make(map[string]*stats),
		r:      proto.RouterServiceClient("go.micro.srv.router", c),
	}

	go o.run()
	return o
}

func (o *os) newStats(s *registry.Service, node *registry.Node) {
	o.Lock()
	defer o.Unlock()

	if _, ok := o.stats[node.Id]; ok {
		return
	}

	o.stats[node.Id] = newStats(&registry.Service{
		Name:     s.Name,
		Version:  s.Version,
		Metadata: s.Metadata,
		Nodes:    []*registry.Node{node},
	})
}

func (o *os) publish() {
	o.RLock()
	defer o.RUnlock()

	opts := o.server.Options()

	// temporarily build client Service
	// should just be pulled from opts.Service()
	// or something like that
	var addr, host string
	var port int

	if len(opts.Advertise) > 0 {
		addr = opts.Advertise
	} else {
		addr = opts.Address
	}

	parts := strings.Split(addr, ":")
	if len(parts) == 2 {
		i, _ := strconv.ParseInt(parts[1], 10, 64)
		host = parts[0]
		port = int(i)
	} else {
		host = addr
	}

	// the client service. this is us
	service := &registry.Service{
		Name:    opts.Name,
		Version: opts.Version,
		Nodes: []*registry.Node{&registry.Node{
			Id:       opts.Id,
			Address:  host,
			Port:     port,
			Metadata: opts.Metadata,
		}},
	}

	// publish all the stats and reset
	for _, stat := range o.stats {
		// create publication
		msg := o.client.NewPublication(StatsTopic, stat.ToProto(service))
		// reset the stats
		stat.Reset()
		// publish message
		go o.client.Publish(context.TODO(), msg)
	}
}

func (o *os) subscribe() {
	// TODO: subscribe to stream of updates from router
	// send request for every service
	// recv for every service
	// create cache
	// create stats

	streams := make(map[string]bool)
	t := time.NewTicker(publishInterval)

	fn := func(service string, exit chan bool) {
		for {
			select {
			case <-exit:
				return
			default:
				o.stream(service)
				time.Sleep(time.Millisecond * 100)
			}
		}
	}

	for {
		select {
		case <-o.exit:
			t.Stop()
			return
		case <-t.C:
			o.RLock()
			cache := o.cache
			o.RUnlock()

			for name, _ := range cache {
				if _, ok := streams[name]; ok {
					continue
				}
				fn(name, o.exit)
				streams[name] = true
			}
		}
	}
}

func (o *os) stream(service string) {
	stream, err := o.r.SelectStream(context.TODO(), &proto.SelectRequest{Service: service})
	if err != nil {
		return
	}

	exit := make(chan bool)

	defer func() {
		close(exit)
	}()

	go func() {
		select {
		case <-exit:
			// probably errored
		case <-o.exit:
			stream.Close()
		}
	}()

	for {
		rsp, err := stream.Recv()
		if err != nil {
			return
		}

		var services []*registry.Service
		nodes := make(map[string]bool)

		for _, serv := range rsp.Services {
			rservice := protoToService(serv)
			services = append(services, rservice)

			// create stats
			for _, node := range rservice.Nodes {
				o.newStats(rservice, node)
				nodes[node.Id] = true
			}
		}

		o.Lock()
		// delete nodes from stats that have been removed
		// we might actually lost stats by doing this
		// TODO: move to a reaper
		if s, ok := o.cache[service]; ok {
			for node, _ := range s.nodes {
				if _, ok := nodes[node]; !ok {
					delete(o.stats, node)
				}
			}
		}

		// cache the service
		o.cache[service] = newCache(services, rsp.Expires)
		o.Unlock()
	}
}

func (o *os) rselect(service string) (*cache, error) {
	// check the cache
	o.RLock()
	if c, ok := o.cache[service]; ok {
		if c.expires == -1 || c.expires > time.Now().Unix() {
			o.RUnlock()
			return c, nil
		}
	}
	o.RUnlock()

	// not cached or expired

	// call router to get selection for service
	rsp, err := o.r.Select(context.TODO(), &proto.SelectRequest{
		Service: service,
	})

	// error then bail
	if err != nil {
		return nil, err
	}

	// translate from proto to *registry.Service
	var services []*registry.Service
	for _, serv := range rsp.Services {
		rservice := protoToService(serv)
		services = append(services, rservice)

		// create stats
		for _, node := range rservice.Nodes {
			o.newStats(rservice, node)
		}
	}

	// cache the service
	c := newCache(services, rsp.Expires)
	o.Lock()
	o.cache[service] = c
	o.Unlock()

	return c, nil
}

func (o *os) run() {
	t := time.NewTicker(publishInterval)

	go o.subscribe()

	for {
		select {
		case <-o.exit:
			t.Stop()
			return
		case <-t.C:
			o.publish()
		}
	}
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

func (o *os) Init(opts ...selector.Option) error {
	var options selector.Options
	for _, o := range opts {
		o(&options)
	}

	// TODO: Fix. This might all be really bad and hacky

	if c, ok := client.FromContext(options.Context); ok {
		o.client = c
		o.r = proto.RouterServiceClient("go.micro.srv.router", c)
	}

	if s, ok := server.FromContext(options.Context); ok {
		o.server = s
	}

	return nil
}

func (o *os) Options() selector.Options {
	return o.opts
}

func (o *os) Record(r Request, node *registry.Node, d time.Duration, err error) {
	o.Lock()
	defer o.Unlock()

	stats, ok := o.stats[node.Id]
	if !ok {
		return
	}

	stats.Record(r, node, d, err)
}

func (o *os) Stats() ([]*Stats, error) {
	return nil, nil
}

func (o *os) Select(service string, opts ...selector.SelectOption) (selector.Next, error) {
	// create options
	var options selector.SelectOptions
	for _, o := range opts {
		o(&options)
	}

	// get service from the cache
	// or call the router for list
	cache, err := o.rselect(service)
	if err != nil {
		return nil, err
	}
	return cache.Filter(options.Filters)
}

func (o *os) Mark(service string, node *registry.Node, err error) {
	o.Lock()
	defer o.Unlock()

	stats, ok := o.stats[node.Id]
	if !ok {
		return
	}

	// mark result for the node
	stats.Mark(service, node, err)
}

func (o *os) Reset(service string) {
	o.Lock()
	defer o.Unlock()

	// reset stats for the service
	for _, stat := range o.stats {
		if stat.service.Name == service {
			stat.Reset()
		}
	}
}

func (o *os) String() string {
	return "os"
}
