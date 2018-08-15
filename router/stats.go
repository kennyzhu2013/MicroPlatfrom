package router

import (
	"time"

	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/registry"
	router "github.com/micro/go-os/router/proto"
	"github.com/rcrowley/go-metrics"
)

type stats struct {
	// the service for these stats
	service *registry.Service

	// time of stats creation
	timestamp int64

	// used for the selector stats
	dropped int64
	success int64
	errors  int64

	// track slas for endpoints
	endpoints map[string]*endpoint
}

// endpoint metrics
type endpoint struct {
	dropped metrics.Timer
	success metrics.Timer
	errors  metrics.Timer
}

func newEndpoint() *endpoint {
	return &endpoint{
		dropped: metrics.NewTimer(),
		success: metrics.NewTimer(),
		errors:  metrics.NewTimer(),
	}
}

func newStats(service *registry.Service) *stats {
	return &stats{
		service:   service,
		endpoints: make(map[string]*endpoint),
		timestamp: time.Now().Unix(),
	}
}

func (s *stats) Mark(service string, node *registry.Node, err error) {
	if err == nil {
		s.success++
		return
	}

	// parse so we can get the error code and see if it was dropped (408)
	perr, ok := err.(*errors.Error)
	if !ok {
		perr = &errors.Error{Id: service, Code: 500, Detail: err.Error()}
	}

	switch perr.Code {
	case 408:
		s.dropped++
	default:
		s.errors++
	}
}

func (s *stats) Record(r Request, node *registry.Node, d time.Duration, err error) {
	endpoint, ok := s.endpoints[r.Method()]
	if !ok {
		endpoint = newEndpoint()
		s.endpoints[r.Method()] = endpoint
	}

	// track the success
	if err == nil {
		endpoint.success.Update(d)
		return
	}

	// track the error

	// parse so we can get the error code and see if it was dropped (408)
	perr, ok := err.(*errors.Error)
	if !ok {
		perr = &errors.Error{Id: r.Service(), Code: 500, Detail: err.Error()}
	}

	switch perr.Code {
	case 408:
		endpoint.dropped.Update(d)
	default:
		endpoint.errors.Update(d)
	}
}

func (s *stats) Reset() {
	s.timestamp = time.Now().Unix()
	s.dropped = 0
	s.success = 0
	s.errors = 0

	for k, _ := range s.endpoints {
		s.endpoints[k] = newEndpoint()
	}
}

func (s *stats) ToProto(client *registry.Service) *router.Stats {
	t := time.Now().Unix()

	stats := &router.Stats{
		Service:   serviceToProto(s.service),
		Client:    serviceToProto(client),
		Timestamp: t,
		Duration:  t - s.timestamp,
		Selected: &router.Selected{
			Errors:  &router.Metrics{Count: s.errors},
			Success: &router.Metrics{Count: s.success},
			Dropped: &router.Metrics{Count: s.dropped},
		},
	}

	for name, endpoint := range s.endpoints {
		stats.EndpointStats = append(stats.EndpointStats, &router.EndpointStats{
			Name: name,
			Errors: &router.Metrics{
				Count:   endpoint.errors.Count(),
				Mean:    endpoint.errors.Mean(),
				StdDev:  endpoint.errors.StdDev(),
				Upper95: endpoint.errors.Percentile(0.95),
			},
			Success: &router.Metrics{
				Count:   endpoint.success.Count(),
				Mean:    endpoint.success.Mean(),
				StdDev:  endpoint.success.StdDev(),
				Upper95: endpoint.success.Percentile(0.95),
			},
			Dropped: &router.Metrics{
				Count:   endpoint.dropped.Count(),
				Mean:    endpoint.dropped.Mean(),
				StdDev:  endpoint.dropped.StdDev(),
				Upper95: endpoint.dropped.Percentile(0.95),
			},
		})
	}

	return stats
}
