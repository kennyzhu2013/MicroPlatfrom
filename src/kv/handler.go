package kv

import (
	"sync"
	"time"

	"github.com/micro/go-micro/errors"
	proto "github.com/micro/go-os/kv/proto"

	"golang.org/x/net/context"
)

type kv struct{}

type item struct {
	*proto.Item
	Timestamp int64
}

var (
	// really shouldn't do this
	// fix later
	mtx   sync.RWMutex
	items = map[string]*item{}
)

func (k *kv) Get(ctx context.Context, req *proto.GetRequest, rsp *proto.GetResponse) error {
	mtx.RLock()
	defer mtx.RUnlock()

	item, ok := items[req.Key]
	if !ok {
		return errors.NotFound(serviceName, "not found")
	}

	// non expiring key
	if item.Expiration <= 0 {
		rsp.Item = item.Item
		return nil
	}

	// expiring key

	t := time.Now().Unix()

	if delta := t - item.Timestamp; delta > item.Expiration {
		// not yet reaped but you can't have it either
		return errors.NotFound(serviceName, "not found")
	}

	rsp.Item = item.Item
	return nil
}

func (k *kv) Put(ctx context.Context, req *proto.PutRequest, rsp *proto.PutResponse) error {
	mtx.Lock()
	items[req.Item.Key] = &item{req.Item, time.Now().Unix()}
	mtx.Unlock()
	return nil
}

func (k *kv) Del(ctx context.Context, req *proto.DelRequest, rsp *proto.DelResponse) error {
	mtx.Lock()
	delete(items, req.Key)
	mtx.Unlock()
	return nil
}
