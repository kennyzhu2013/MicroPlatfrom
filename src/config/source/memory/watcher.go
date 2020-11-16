package memory

import (
	"github.com/micro/go-os/config"
)

type Watcher struct {
	Id      string
	Updates chan *config.ChangeSet
	Source  *Source
}

func (w *Watcher) Next() (*config.ChangeSet, error) {
	cs := <-w.Updates
	return cs, nil
}

func (w *Watcher) Stop() error {
	w.Source.Lock()
	delete(w.Source.Watchers, w.Id)
	w.Source.Unlock()
	return nil
}
