package monitor

import (
	"syscall"
	"time"

	"github.com/micro/go-micro/errors"
	proto "github.com/micro/go-os/monitor/proto"
)

type stats struct {
	utime   syscall.Timeval
	stime   syscall.Timeval
	maxRss  int64
	inBlock int64
	ouBlock int64
	volCtx  int64
	invCtx  int64
	timers  map[string]*endpointTimers
}

type endpointTimers struct {
	service string
	method  string
	dropped metrics.Timer
	success metrics.Timer
	errors  metrics.Timer
}

func newStats() *stats {
	var r syscall.Rusage

	if err := syscall.Getrusage(0, &r); err != nil {
		return nil
	}

	return &stats{
		utime:   r.Utime,
		stime:   r.Stime,
		maxRss:  r.Maxrss,
		inBlock: r.Inblock,
		ouBlock: r.Oublock,
		volCtx:  r.Nvcsw,
		invCtx:  r.Nivcsw,
		timers:  make(map[string]*endpointTimers),
	}
}

func (s *stats) endpoints() []*proto.Endpoint {
	var endpoints []*proto.Endpoint

	for _, timer := range s.timers {
		endpoints = append(endpoints, &proto.Endpoint{
			Service: timer.service,
			Method:  timer.method,
			Errors: &proto.Metrics{
				Count:   timer.errors.Count(),
				Mean:    timer.errors.Mean(),
				StdDev:  timer.errors.StdDev(),
				Upper95: timer.errors.Percentile(0.95),
			},
			Success: &proto.Metrics{
				Count:   timer.success.Count(),
				Mean:    timer.success.Mean(),
				StdDev:  timer.success.StdDev(),
				Upper95: timer.success.Percentile(0.95),
			},
			Dropped: &proto.Metrics{
				Count:   timer.dropped.Count(),
				Mean:    timer.dropped.Mean(),
				StdDev:  timer.dropped.StdDev(),
				Upper95: timer.dropped.Percentile(0.95),
			},
		})
	}

	return endpoints
}

func (s *stats) record(r Request, d time.Duration, err error) {
	t, ok := s.timers[Service()+Method()]
	if !ok {
		t = &endpointTimers{
			service: Service(),
			method:  Method(),
			dropped: metrics.NewTimer(),
			success: metrics.NewTimer(),
			errors:  metrics.NewTimer(),
		}
		s.timers[Service()+Method()] = t
	}

	// no error, mark success and return
	if err == nil {
		t.success.Update(d)
		return
	}

	// parse the error
	perr, ok := err.(*errors.Error)
	if !ok {
		perr = &errors.Error{Code: 500}
	}

	// check if its a timeout
	switch perr.Code {
	case 408:
		t.dropped.Update(d)
	default:
		t.errors.Update(d)
	}
}
