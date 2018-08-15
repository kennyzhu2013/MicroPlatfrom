package discovery

import (
	"sync"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"

	proto2 "github.com/micro/discovery-srv/proto/registry"
	proto "github.com/micro/go-os/discovery/proto"
	"golang.org/x/net/context"
)

type os struct {
	exit chan bool

	options registry.Options
	opts    *Options
	next    func() []string

	sync.RWMutex
	heartbeats map[string]*proto.Heartbeat
	cache      map[string][]*registry.Service
}

func newOS(opts ...registry.Option) Discovery {
	options := registry.Options{
		Context: context.Background(),
	}

	for _, o := range opts {
		o(&options)
	}

	dopts := getOptions(options.Context)

	// set default client
	if dopts.Client == nil {
		dopts.Client = client.DefaultClient
	}

	// set default intervale
	if dopts.Interval == time.Duration(0) {
		dopts.Interval = time.Second * 30
	}

	o := &os{
		options:    options,
		opts:       dopts,
		next:       rr(options.Addrs),
		exit:       make(chan bool),
		heartbeats: make(map[string]*proto.Heartbeat),
		cache:      make(map[string][]*registry.Service),
	}

	go o.run()
	return o
}

func (o *os) heartbeat() {
	t := time.NewTicker(o.opts.Interval)

	for {
		select {
		case <-t.C:
			var heartbeats []*proto.Heartbeat

			o.RLock()
			for _, hb := range o.heartbeats {
				heartbeats = append(heartbeats, hb)
			}
			o.RUnlock()

			for _, hb := range heartbeats {
				hb.Timestamp = time.Now().Unix()
				pub := o.opts.Client.NewPublication(HeartbeatTopic, hb)
				o.opts.Client.Publish(context.TODO(), pub)
			}
		case <-o.exit:
			return
		}
	}
}

func (o *os) watch(ch chan *registry.Result) {
	watch, err := o.Watch()
	for {
		if err == nil {
			break
		}
		select {
		case <-o.exit:
			return
		default:
			time.Sleep(time.Second)
			watch, err = o.Watch()
		}
	}
	defer watch.Stop()

	for {
		next, err := watch.Next()
		if err != nil {
			w, err := o.Watch()
			if err != nil {
				time.Sleep(time.Second)
				continue
			}
			watch.Stop()
			watch = w
			time.Sleep(time.Second)
			continue
		}
		select {
		case ch <- next:
		case <-o.exit:
			return
		}
	}
}

func (o *os) run() {
	ch := make(chan *registry.Result)

	go o.watch(ch)
	go o.heartbeat()

	for {
		select {
		case <-o.exit:
			return
		case next, ok := <-ch:
			if !ok {
				return
			}
			o.update(next)
		}
	}
}

func (o *os) update(res *registry.Result) {
	if res == nil || res.Service == nil {
		return
	}

	o.Lock()
	defer o.Unlock()

	services, ok := o.cache[res.Service.Name]
	if !ok {
		// we're not going to cache anything
		// unless there was already a lookup
		return
	}

	if len(res.Service.Nodes) == 0 {
		switch res.Action {
		case "delete":
			delete(o.cache, res.Service.Name)
		}
		return
	}

	// existing service found
	var service *registry.Service
	var index int
	for i, s := range services {
		if s.Version == res.Service.Version {
			service = s
			index = i
		}
	}

	switch res.Action {
	case "create", "update":
		if service == nil {
			services = append(services, res.Service)
			o.cache[res.Service.Name] = services
			return
		}

		// append old nodes to new service
		for _, cur := range service.Nodes {
			var seen bool
			for _, node := range res.Service.Nodes {
				if cur.Id == node.Id {
					seen = true
					break
				}
			}
			if !seen {
				res.Service.Nodes = append(res.Service.Nodes, cur)
			}
		}

		services[index] = res.Service
		o.cache[res.Service.Name] = services
	case "delete":
		if service == nil {
			return
		}

		var nodes []*registry.Node

		// filter cur nodes to remove the dead one
		for _, cur := range service.Nodes {
			var seen bool
			for _, del := range res.Service.Nodes {
				if del.Id == cur.Id {
					seen = true
					break
				}
			}
			if !seen {
				nodes = append(nodes, cur)
			}
		}

		if len(nodes) == 0 {
			if len(services) == 1 {
				delete(o.cache, service.Name)
			} else {
				var srvs []*registry.Service
				for _, s := range services {
					if s.Version != service.Version {
						srvs = append(srvs, s)
					}
				}
				o.cache[service.Name] = srvs
			}
			return
		}

		service.Nodes = nodes
		services[index] = service
		o.cache[res.Service.Name] = services
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

func (o *os) Register(s *registry.Service, opts ...registry.RegisterOption) error {
	var grr error
	service := toProto(s)
	req := o.opts.Client.NewRequest(
		"go.micro.srv.discovery",
		"Registry.Register",
		&proto2.RegisterRequest{
			Service: service,
		},
	)

	for _, addr := range o.next() {
		rsp := &proto2.RegisterResponse{}
		err := o.opts.Client.CallRemote(context.TODO(), addr, req, rsp)
		if err != nil {
			grr = err
		}
	}

	// create a heartbeat for this service
	o.Lock()
	hb := &proto.Heartbeat{
		Id:       s.Nodes[0].Id,
		Service:  service,
		Interval: int64(o.opts.Interval.Seconds()),
		Ttl:      int64((o.opts.Interval.Seconds()) * 1.25),
	}
	o.heartbeats[hb.Id] = hb
	o.Unlock()

	return grr
}

func (o *os) Deregister(s *registry.Service) error {
	var grr error
	service := toProto(s)
	req := o.opts.Client.NewRequest(
		"go.micro.srv.discovery",
		"Registry.Deregister",
		&proto2.DeregisterRequest{
			Service: service,
		},
	)

	for _, addr := range o.next() {
		rsp := &proto2.RegisterResponse{}
		err := o.opts.Client.CallRemote(context.TODO(), addr, req, rsp)
		if err != nil {
			grr = err
		}
	}

	// remove heartbeat
	o.Lock()
	delete(o.heartbeats, s.Nodes[0].Id)
	o.Unlock()

	return grr
}

func (o *os) GetService(name string) ([]*registry.Service, error) {
	o.RLock()
	if services, ok := o.cache[name]; ok {
		o.RUnlock()
		return services, nil
	}
	o.RUnlock()

	var grsp *proto2.GetServiceResponse
	var grr error

	req := o.opts.Client.NewRequest(
		"go.micro.srv.discovery",
		"Registry.GetService",
		&proto2.GetServiceRequest{
			Service: name,
		},
	)

	for _, addr := range o.next() {
		rsp := &proto2.GetServiceResponse{}
		err := o.opts.Client.CallRemote(context.TODO(), addr, req, rsp)
		if err != nil {
			grr = err
			continue
		}
		grsp = rsp
		break
	}

	if grr != nil {
		return nil, grr
	}

	var services []*registry.Service
	for _, service := range grsp.Services {
		services = append(services, toService(service))
	}

	// cache on lookup
	o.Lock()
	o.cache[name] = services
	o.Unlock()
	return services, nil
}

// TODO: prepopulate the cache
func (o *os) ListServices() ([]*registry.Service, error) {
	o.RLock()
	if cache := o.cache; len(cache) > 0 {
		var services []*registry.Service
		for _, service := range cache {
			services = append(services, service...)
		}
		o.RUnlock()
		return services, nil
	}
	o.RUnlock()

	var grsp *proto2.ListServicesResponse
	var grr error

	req := o.opts.Client.NewRequest("go.micro.srv.discovery", "Registry.ListServices", &proto2.ListServicesRequest{})

	for _, addr := range o.next() {
		rsp := &proto2.ListServicesResponse{}
		err := o.opts.Client.CallRemote(context.TODO(), addr, req, rsp)
		if err != nil {
			grr = err
			continue
		}
		grsp = rsp
		break
	}

	if grr != nil {
		return nil, grr
	}

	var services []*registry.Service
	for _, service := range grsp.Services {
		services = append(services, toService(service))
	}
	return services, nil
}

// TODO: subscribe to events rather than the registry itself?
func (o *os) Watch(opts ...registry.WatchOption) (registry.Watcher, error) {
	var wo registry.WatchOptions
	for _, o := range opts {
		o(&wo)
	}

	req := o.opts.Client.NewRequest("go.micro.srv.discovery", "Registry.Watch", &proto2.WatchRequest{
		Service: wo.Service,
	})

	var gstream client.Streamer
	var grr error

	for _, addr := range o.next() {
		stream, err := o.opts.Client.StreamRemote(context.TODO(), addr, req)
		if err != nil {
			grr = err
			continue
		}
		gstream = stream
		break
	}

	if grr != nil {
		return nil, grr
	}

	// send empty watch request
	if err := gstream.Send(&proto2.WatchRequest{}); err != nil {
		return nil, err
	}

	wc := &watchClient{gstream}
	return &watcher{wc}, nil
}

func (o *os) Options() registry.Options {
	return o.options
}

func (o *os) String() string {
	return "os"
}
