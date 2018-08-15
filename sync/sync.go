// Package sync is an interface for synchronization.
package sync

import (
	"github.com/micro/go-micro/registry"
)

const (
	FollowerStatus  LeaderStatus = 0
	CandidateStatus LeaderStatus = 1
	ElectedStatus   LeaderStatus = 2
)

type LeaderStatus int32

type Sync interface {
	// distributed lock interface
	Lock(id string, opts ...LockOption) (Lock, error)
	// leader election interface
	Leader(id string, opts ...LeaderOption) (Leader, error)
	// Name of sync
	String() string
}

type Lock interface {
	// The unique ID to lock on
	Id() string
	// Acquire the lock
	Acquire() error
	// Release the lock
	Release() error
}

type Leader interface {
	// The unique ID to synchronize with
	Id() string
	// Returns the current leader
	Leader() (*registry.Node, error)
	// Elect self to become leader
	Elect() (Elected, error)
	// Returns the status of this node
	Status() (LeaderStatus, error)
}

type Elected interface {
	// Returns a channel which indicates
	// when the leadership is revoked
	Revoked() (chan struct{}, error)
	// Resign the leadership
	Resign() error
}

type Option func(o *Options)

type LockOption func(o *LockOptions)

type LeaderOption func(o *LeaderOptions)

var (
	DefaultNamespace = "/micro/sync"
)
