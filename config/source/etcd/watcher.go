package etcd

import (
	"crypto/md5"
	"errors"
	"fmt"

	"github.com/coreos/etcd/client"
	"github.com/micro/go-os/config"

	"golang.org/x/net/context"
)

type watcher struct {
	name string

	w      client.Watcher
	ctx    context.Context
	cancel func()
	exit   chan bool
}

func newWatcher(key string, addrs []string, name string) (config.SourceWatcher, error) {
	c, err := client.New(client.Config{
		Endpoints: addrs,
	})
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	w := client.NewKeysAPI(c).Watcher(key, nil)

	return &watcher{
		name:   name,
		ctx:    ctx,
		cancel: cancel,
		exit:   make(chan bool),
		w:      w,
	}, nil
}

func (w *watcher) Next() (*config.ChangeSet, error) {
	select {
	case <-w.exit:
		return nil, errors.New("watcher stopped")
	default:
		rsp, err := w.w.Next(w.ctx)
		if err != nil {
			return nil, err
		}

		// hash the etcd
		h := md5.New()
		h.Write([]byte(rsp.Node.Value))
		checksum := fmt.Sprintf("%x", h.Sum(nil))

		return &config.ChangeSet{
			Source:   w.name,
			Data:     []byte(rsp.Node.Value),
			Checksum: checksum,
		}, nil
	}
}

func (w *watcher) Stop() error {
	select {
	case <-w.exit:
		return nil
	default:
		close(w.exit)
		w.cancel()
	}
	return nil
}
