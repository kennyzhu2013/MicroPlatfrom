package kv

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	proto "github.com/micro/go-os/kv/proto"
	store "github.com/micro/kv-srv/proto/store"

	"github.com/asim/consistent"
	"golang.org/x/net/context"
)

/*
	OS KV is a consistently hashed in memory key-value store utilising
	all the services in the network. Aww yea. Can optionally be namespaced using that provided

	Optional the kv-srv can be used rather than having each client participate in the ring.
*/

type os struct {
	opts Options
	exit chan bool
	hash *consistent.Consistent

	// client
	client store.StoreClient

	sync.RWMutex
	nodes map[string]int64
}

type Announcement struct {
	// Which namespace does it belong to
	Namespace string
	Address   string
	Timestamp int64
}

var (
	// not really needed but you know...
	serviceName = "go.micro.srv.kv"

	GossipTopic = "go.micro.kv.announce"
	GossipEvent = time.Second * 1
	ReaperEvent = time.Second * 10
)

func newOS(opts ...Option) KV {
	options := Options{
		Internal: true,
	}
	for _, o := range opts {
		o(&options)
	}

	if options.Client == nil {
		options.Client = client.DefaultClient
	}

	if options.Server == nil {
		options.Server = server.DefaultServer
	}

	if options.Replicas == 0 {
		options.Replicas = 1
	}

	o := &os{
		exit:   make(chan bool),
		opts:   options,
		hash:   consistent.New(),
		nodes:  make(map[string]int64),
		client: store.NewStoreClient(serviceName, options.Client),
	}

	// If using gossip then add the handlers and run the broadcaster
	if !o.opts.Service {
		options.Server.Subscribe(
			options.Server.NewSubscriber(
				GossipTopic,
				o.subscriber,
				server.InternalSubscriber(options.Internal),
			),
		)

		options.Server.Handle(
			options.Server.NewHandler(
				&proto.KV{new(kv)},
				server.InternalHandler(options.Internal),
			),
		)

		go o.run()
	}

	return o
}

func (a *Announcement) Topic() string {
	return GossipTopic
}

func (a *Announcement) Message() interface{} {
	return a
}

func (a *Announcement) ContentType() string {
	return "application/json"
}

func (o *os) address() string {
	config := o.opts.Server.Options()

	var advt, host string
	var port int

	if len(config.Advertise) > 0 {
		advt = config.Advertise
	} else {
		advt = config.Address
	}

	parts := strings.Split(advt, ":")
	if len(parts) > 1 {
		host = strings.Join(parts[:len(parts)-1], ":")
		port, _ = strconv.Atoi(parts[len(parts)-1])
	} else {
		host = parts[0]
	}

	addr, _ := extractAddress(host)

	if port > 0 {
		return fmt.Sprintf("%s:%d", addr, port)
	}

	return addr
}

func (o *os) reap() {
	t := time.Now().Unix()
	r := int64(GossipEvent.Seconds() * 1.5)

	// reap nodes
	o.Lock()
	for node, seen := range o.nodes {
		// Is last greater than GossipEvent time plus some
		if last := t - seen; last > r {
			delete(o.nodes, node)
			o.hash.Remove(node)
		}
	}
	o.Unlock()

	// reap keys
	mtx.Lock()

	for key, item := range items {
		// don't expire zero or less
		if item.Expiration <= 0 {
			continue
		}

		// Delta greater than expiration
		if delta := t - item.Timestamp; delta > item.Expiration {
			delete(items, key)
		}
	}
	mtx.Unlock()
}

func (o *os) run() {
	// setup the ticker
	t := time.NewTicker(GossipEvent)
	r := time.NewTicker(ReaperEvent)

	o.setup()

	// now lets go!
	for {
		select {
		case <-t.C:
			o.publish()
		case <-r.C:
			o.reap()
		case <-o.exit:
			t.Stop()
			r.Stop()
			return
		}
	}
}

func (o *os) publish() error {
	a := &Announcement{
		Namespace: o.opts.Namespace,
		Address:   o.address(),
		Timestamp: time.Now().Unix(),
	}
	return o.opts.Client.Publish(context.TODO(), a)
}

// immediately add self to ring
func (o *os) setup() {
	for i := 0; i < 10; i++ {
		// wait till there's a valid address from the server
		if p := strings.Split(o.address(), ":"); len(p) < 2 {
			time.Sleep(GossipEvent / 100.0)
			continue
		}
		// have a valid address, setup, now
		o.subscriber(context.Background(), &Announcement{
			Namespace: o.opts.Namespace,
			Address:   o.address(),
			Timestamp: time.Now().Unix(),
		})
		return
	}
}

func (o *os) subscriber(ctx context.Context, a *Announcement) error {
	o.Lock()
	defer o.Unlock()

	if o.opts.Namespace != a.Namespace {
		return nil
	}

	_, ok := o.nodes[a.Address]
	if !ok {
		o.hash.Add(a.Address)
	}

	o.nodes[a.Address] = a.Timestamp
	return nil
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

func (o *os) Get(key string) (*Item, error) {
	// if we're using the KV service then call that
	if o.opts.Service {
		rsp, err := o.client.Get(context.TODO(), &store.GetRequest{
			Key: key,
		})
		if err != nil {
			return nil, err
		}
		return &Item{
			Key:        rsp.Item.Key,
			Value:      rsp.Item.Value,
			Expiration: time.Duration(rsp.Item.Expiration) * time.Second,
		}, nil
	}

	nodes, err := o.hash.GetN(key, o.opts.Replicas)
	if err != nil {
		return nil, err
	}

	req := o.opts.Client.NewRequest(serviceName, "KV.Get", &proto.GetRequest{
		Key: key,
	})

	rsp := &proto.GetResponse{}

	for _, node := range nodes {
		// query node and return
		if err := o.opts.Client.CallRemote(context.TODO(), node, req, rsp); err == nil {
			if rsp.Item == nil {
				continue
			}
			return &Item{
				Key:        rsp.Item.Key,
				Value:      rsp.Item.Value,
				Expiration: time.Second * time.Duration(rsp.Item.Expiration),
			}, nil
		}
	}

	return nil, ErrNotFound
}

func (o *os) Del(key string) error {
	// if we're using the KV service then call that
	if o.opts.Service {
		_, err := o.client.Del(context.TODO(), &store.DelRequest{
			Key: key,
		})
		return err
	}

	nodes, err := o.hash.GetN(key, o.opts.Replicas)
	if err != nil {
		return err
	}

	req := o.opts.Client.NewRequest(serviceName, "KV.Del", &proto.DelRequest{
		Key: key,
	})

	var gerr error

	for _, node := range nodes {
		rsp := &proto.DelResponse{}
		if err := o.opts.Client.CallRemote(context.TODO(), node, req, rsp); err != nil {
			gerr = err
		}
	}

	return gerr
}

func (o *os) Put(item *Item) error {
	// if we're using the KV service then call that
	if o.opts.Service {
		_, err := o.client.Put(context.TODO(), &store.PutRequest{
			Item: &store.Item{
				Key:        item.Key,
				Value:      item.Value,
				Expiration: int64(item.Expiration.Seconds()),
			},
		})
		return err
	}

	nodes, err := o.hash.GetN(item.Key, o.opts.Replicas)
	if err != nil {
		return err
	}

	req := o.opts.Client.NewRequest(serviceName, "KV.Put", &proto.PutRequest{
		Item: &proto.Item{
			Key:        item.Key,
			Value:      item.Value,
			Expiration: int64(item.Expiration.Seconds()),
		},
	})

	var gerr error

	for _, node := range nodes {
		rsp := &proto.PutResponse{}
		if err := o.opts.Client.CallRemote(context.TODO(), node, req, rsp); err != nil {
			gerr = err
		}
	}

	return gerr
}

func (o *os) String() string {
	return "os"
}
