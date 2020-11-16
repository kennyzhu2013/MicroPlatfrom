# Monitor [![GoDoc](https://godoc.org/github.com/micro/go-os?status.svg)](https://godoc.org/github.com/micro/go-os/monitor)

Provides a high level pluggable abstraction for monitoring. 

<p align="center">
  <img src="https://github.com/micro/go-os/blob/master/doc/monitor.png" />
</p>

## Interface

Allows the ability for user defined healthchecks. Exactly what's required for monitoring 
business logic as opposed to lower level things like pings, tcp checks, etc. Also 
collates stats about the service and periodically publishes.

```go
// The monitor aggregates, observes and publishes
// information about the current process.
// This includes status; started/running/stopped,
// stats; cpu, memory, runtime and healthchecks.
type Monitor interface {
	Close() error
	Checker
	Stats
	String() string
}

// Checker interface allows creation of healthchecks
type Checker interface {
	NewHealthChecker(id, desc string, fn HealthCheck) HealthChecker
	Register(HealthChecker) error
	Deregister(HealthChecker) error
	HealthChecks() ([]HealthChecker, error)
}

// represents a healthcheck function
type HealthChecker interface {
	// Unique id of the healthcheck
	Id() string
	// Description of what the healthcheck does
	Description() string
	// Runs the the healthcheck
	Run() (map[string]string, error)
	// Returns the status of the last run
	// Better where the healthcheck is expensive to run
	Status() (map[string]string, error)
}

// stats interface allows recording of endpoint stats
type Stats interface {
	RecordStat(r Request, d time.Duration, err error)
}

type HealthCheck func() (map[string]string, error)

func NewMonitor(opts ...Option) Monitor {
	return newOS(opts...)
}
```

## Supported Backends

- [Monitor service](https://github.com/micro/monitor-srv)

## Usage

```go

import (
	"errors"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-os/monitor"
)

...

m := monitor.NewMonitor(
	// publish intervale
	monitor.Interval(time.Second * 10),
)
defer m.Close()

hc := m.NewHealthChecker(
	// check id
	"com.example.check.myproduct",
	// description
	"Checking product status",
	// healthcheck function
	func() (map[string]string, error) {
		// check some business log
		stats := ProductMetrics()

		if stats.Errors > 10 {
			return nil, errors.New("Product is failing")
		}

		return map[string]string{
			"stats": stats.Metrics(),
			"info":  stats.Info(),
		}, nil
	},
)

m.Register(hc)

// Additionally use client and handler wrappers to let the monitoring service keep track of endpoint stats.

service := micro.NewService(
	micro.Name("com.example.srv.service"),
	micro.WrapClient(monitor.ClientWrapper(m)),
	micro.WrapHandler(monitor.HandlerWrapper(m)),
)
```
