# Source

Source is an interface that defines a backend source from which config is loaded.

```go
// Source is the source from which config is loaded.
// This may be a file, a url, consul, etc.
type Source interface {
	// Loads ChangeSet from the source
	Read() (*ChangeSet, error)
	// Watch for source changes
	// Returns the entire changeset
	Watch() (SourceWatcher, error)
	// Name of source
	String() string
}

// SourceWatcher allows you to watch a source for changes
// Next is a blocking call which returns the next
// ChangeSet update. Stop Renders the watcher unusable.
type SourceWatcher interface {
	Next() (*ChangeSet, error)
	Stop() error
}

// ChangeSet represents a set an actual source
type ChangeSet struct {
	// The time at which the last change occured
	Timestamp time.Time
	// The raw data set for the change
	Data []byte
	// Hash of the source data
	Checksum string
	// The source of this change; file, consul, etcd
	Source string
}
```
