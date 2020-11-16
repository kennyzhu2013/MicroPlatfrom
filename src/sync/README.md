# Sync [![GoDoc](https://godoc.org/github.com/micro/go-os?status.svg)](https://godoc.org/github.com/micro/go-os/sync)

Sync is client library for locking and leadership election. It provides a way to coordinate across a number of nodes. 
It's a building block for synchronization in distributed systems.

Most distributed systems choose AP - availability and partition tolerance. At some point we need consistency as well 
so Sync provides a way to do this. 

```go
type Sync interface {
        // distributed lock interface
        Lock(...LockOption) (Lock, error)
        // leader election interface
        Leader(...LeaderOption) (Leader, error)
}

type Lock interface {
        Id() string
        Acquire() error
        Release() error
}

type Leader interface {
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
        Revoked() (chan bool, error)
        // Resign the leadership
        Resign() error
}
```

## Supported Backends

- [Consul](https://github.com/micro/go-plugins/tree/master/sync/consul)
- [Etcd](https://github.com/micro/go-plugins/tree/master/sync/etcd)

## Usage 

```go
package main

import (
        "fmt"

        "github.com/micro/go-os/sync/consul"
)

func main() {
        c := consul.NewSync()

        // create a lock
        lock, err := c.Lock("lockid")
        if err != nil {
                fmt.Println("Failed to create lock", err)
                return
        }

        fmt.Println("Created lock")

        // acquire the lock
        if err := lock.Acquire(); err != nil {
                fmt.Println("Failed to acquire lock", err)
                return
        }

        fmt.Println("Acquired lock")

        // release the lock
        if err := lock.Release(); err != nil {
                fmt.Println("Failed to release lock", err)
                return
        }

        fmt.Println("Released lock")
}
```
