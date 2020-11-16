// Package monitor is an interface for monitoring.
package monitor

import (
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
)

// The monitor aggregates, observes and publishes
// information about the current process.
// This includes status; started/running/stopped,
// stats; cpu, memory, runtime and healthchecks.
type Monitor interface {
	Checker
	Stats
	Close() error
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

// could be client or server request
type Request interface {
	Service() string
	Method() string
}

type HealthCheck func() (map[string]string, error)

type Option func(*Options)

type Options struct {
	Interval time.Duration
	Client   client.Client
	Server   server.Server
}

var (
	HealthCheckTopic = "micro.monitor.healthcheck"
	StatusTopic      = "micro.monitor.status"
	StatsTopic       = "micro.monitor.stats"
)

func ClientWrapper(m Monitor) client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{c, m}
	}
}

func HandlerWrapper(m Monitor) server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		return handlerWrapper(h, m)
	}
}

func NewMonitor(opts ...Option) Monitor {
	return newOS(opts...)
}
