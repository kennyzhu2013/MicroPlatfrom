package monitor

import (
	"errors"
	"github.com/kennyzhu/go-os/src/monitor/proto"
	"runtime"
	"sync"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"

	"golang.org/x/net/context"
)

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotAvailable  = errors.New("not available")
)

type os struct {
	exit              chan bool
	opts              Options
	name, version, id string

	sync.Mutex
	hc   map[string]HealthChecker
	stat *stats
}

func newOS(opts ...Option) Monitor {
	var opt Options
	for _, o := range opts {
		o(&opt)
	}

	if opt.Client == nil {
		opt.Client = client.DefaultClient
	}

	if opt.Server == nil {
		opt.Server = server.DefaultServer
	}

	if opt.Interval == time.Duration(0) {
		opt.Interval = time.Minute
	}

	c := opt.Server.Options()

	o := &os{
		stat:    newStats(),
		name:    c.Name,
		version: c.Version,
		id:      c.Id,
		opts:    opt,
		exit:    make(chan bool),
		hc:      make(map[string]HealthChecker),
	}

	go o.run()
	return o
}

func (o *os) stats() {
	ostat := o.stat
	s := newStats()

	// update
	o.stat = s

	cpu := &go_micro_os_monitor.CPU{
		UserTime:     uint64(s.utime.Nano() - ostat.utime.Nano()),
		SystemTime:   uint64(s.stime.Nano() - ostat.stime.Nano()),
		VolCtxSwitch: uint64(s.volCtx - ostat.volCtx),
		InvCtxSwitch: uint64(s.invCtx - ostat.invCtx),
	}

	memory := &go_micro_os_monitor.Memory{
		MaxRss: uint64(s.maxRss),
	}

	disk := &go_micro_os_monitor.Disk{
		InBlock: uint64(s.inBlock - ostat.inBlock),
		OuBlock: uint64(s.ouBlock - ostat.ouBlock),
	}

	rm := runtime.MemStats{}
	runtime.ReadMemStats(&rm)

	rtime := &go_micro_os_monitor.Runtime{
		NumThreads: uint64(runtime.NumGoroutine()),
		HeapTotal:  rm.HeapAlloc,
		HeapInUse:  rm.HeapInuse,
	}

	statsProto := &go_micro_os_monitor.Stats{
		Service: &go_micro_os_monitor.Service{
			Name:    o.name,
			Version: o.version,
			Nodes: []*go_micro_os_monitor.Node{&go_micro_os_monitor.Node{
				Id: o.id,
			}},
		},
		Interval:  int64(o.opts.Interval.Seconds()),
		Timestamp: time.Now().Unix(),
		Ttl:       3600,
		Cpu:       cpu,
		Memory:    memory,
		Disk:      disk,
		Runtime:   rtime,
		Endpoints: ostat.endpoints(),
	}

	req := o.opts.Client.NewPublication(StatsTopic, statsProto)
	o.opts.Client.Publish(context.TODO(), req)
}

func (o *os) status(status go_micro_os_monitor.Status_Status) {
	statusProto := &go_micro_os_monitor.Status{
		Status: status,
		Service: &go_micro_os_monitor.Service{
			Name:    o.name,
			Version: o.version,
			Nodes: []*go_micro_os_monitor.Node{&go_micro_os_monitor.Node{
				Id: o.id,
			}},
		},
		Interval:  int64(o.opts.Interval.Seconds()),
		Timestamp: time.Now().Unix(),
		Ttl:       3600,
	}

	req := o.opts.Client.NewPublication(StatusTopic, statusProto)
	o.opts.Client.Publish(context.TODO(), req)
}

func (o *os) update(h HealthChecker) {
	res, err := Run()
	status := go_micro_os_monitor.HealthCheck_OK
	errDesc := ""
	if err != nil {
		status = go_micro_os_monitor.HealthCheck_ERROR
		errDesc = err.Error()
	}

	hcProto := &go_micro_os_monitor.HealthCheck{
		Id:          Id(),
		Description: Description(),
		Timestamp:   time.Now().Unix(),
		Service: &go_micro_os_monitor.Service{
			Name:    o.name,
			Version: o.version,
			Nodes: []*go_micro_os_monitor.Node{&go_micro_os_monitor.Node{
				Id: o.id,
			}},
		},
		Interval: int64(o.opts.Interval.Seconds()),
		Ttl:      3600,
		Status:   status,
		Results:  res,
		Error:    errDesc,
	}

	req := o.opts.Client.NewPublication(HealthCheckTopic, hcProto)
	o.opts.Client.Publish(context.TODO(), req)
}

func (o *os) run() {
	// publish started status
	o.status(go_micro_os_monitor.Status_STARTED)

	t := time.NewTicker(o.opts.Interval)

	for {
		select {
		case <-t.C:
			o.Lock()
			// publish stats
			o.stats()
			// publish status
			o.status(go_micro_os_monitor.Status_RUNNING)
			// publish healthchecks
			for _, check := range o.hc {
				go o.update(check)
			}
			o.Unlock()
		case <-o.exit:
			// stop the ticker
			t.Stop()
			// publish started status
			o.status(go_micro_os_monitor.Status_STOPPED)
			return
		}
	}
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

func (o *os) NewHealthChecker(id, desc string, hc HealthCheck) HealthChecker {
	return newHealthChecker(id, desc, hc)
}

func (o *os) Register(hc HealthChecker) error {
	o.Lock()
	defer o.Unlock()
	if _, ok := o.hc[Id()]; ok {
		return ErrAlreadyExists
	}
	o.hc[Id()] = hc
	return nil
}

func (o *os) Deregister(hc HealthChecker) error {
	o.Lock()
	defer o.Unlock()
	delete(o.hc, Id())
	return nil
}

func (o *os) HealthChecks() ([]HealthChecker, error) {
	var hcs []HealthChecker
	o.Lock()
	for _, hc := range o.hc {
		hcs = append(hcs, hc)
	}
	o.Unlock()
	return hcs, nil
}

func (o *os) RecordStat(r Request, d time.Duration, err error) {
	o.Lock()
	o.stat.record(r, d, err)
	o.Unlock()
}

func (o *os) String() string {
	return "os"
}
