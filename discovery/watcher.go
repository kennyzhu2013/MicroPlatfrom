package discovery

import (
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"

	proto2 "github.com/micro/discovery-srv/proto/registry"
)

type watcher struct {
	wc *watchClient
}

type watchClient struct {
	stream client.Streamer
}

func (w *watcher) Next() (*registry.Result, error) {
	r, err := w.wc.Recv()
	if err != nil {
		return nil, err
	}

	return &registry.Result{
		Action:  r.Result.Action,
		Service: toService(r.Result.Service),
	}, nil
}

func (w *watcher) Stop() {
	w.wc.Close()
}

func (x *watchClient) Close() error {
	return x.stream.Close()
}

func (x *watchClient) Recv() (*proto2.WatchResponse, error) {
	m := new(proto2.WatchResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
