package trace

import (
	"errors"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	proto "github.com/micro/go-os/trace/proto"

	"github.com/pborman/uuid"
	"golang.org/x/net/context"
)

type os struct {
	opts  Options
	spans chan *Span
	exit  chan bool
}

func newOS(opts ...Option) Trace {
	var opt Options
	for _, o := range opts {
		o(&opt)
	}

	if opt.BatchSize == 0 {
		opt.BatchSize = DefaultBatchSize
	}

	if opt.BatchInterval == time.Duration(0) {
		opt.BatchInterval = DefaultBatchInterval
	}

	if len(opt.Topic) == 0 {
		opt.Topic = TraceTopic
	}

	if opt.Client == nil {
		opt.Client = client.DefaultClient
	}

	o := &os{
		exit:  make(chan bool),
		opts:  opt,
		spans: make(chan *Span, 100),
	}

	go o.run()

	return o
}

func serviceToProto(s *registry.Service) *proto.Service {
	if s == nil {
		return nil
	}

	var nodes []*proto.Node

	// add node if it exists
	if len(s.Nodes) > 0 {
		nodes = []*proto.Node{&proto.Node{
			Id:       s.Nodes[0].Id,
			Address:  s.Nodes[0].Address,
			Port:     int64(s.Nodes[0].Port),
			Metadata: s.Nodes[0].Metadata,
		}}
	}

	return &proto.Service{
		Name:     s.Name,
		Version:  s.Version,
		Metadata: s.Metadata,
		Nodes:    nodes,
	}
}

func toProto(s *Span) *proto.Span {
	var annotations []*proto.Annotation

	for _, a := range s.Annotations {
		var timestamp int64
		if !a.Timestamp.IsZero() {
			timestamp = a.Timestamp.UnixNano() / 1e3
		}
		annotations = append(annotations, &proto.Annotation{
			Timestamp: timestamp,
			Type:      proto.Annotation_Type(a.Type),
			Key:       a.Key,
			Value:     a.Value,
			Debug:     a.Debug,
			Service:   serviceToProto(a.Service),
		})
	}

	var timestamp int64
	if !s.Timestamp.IsZero() {
		timestamp = s.Timestamp.UnixNano() / 1e3
	}

	return &proto.Span{
		Name:        s.Name,
		Id:          s.Id,
		TraceId:     s.TraceId,
		ParentId:    s.ParentId,
		Timestamp:   timestamp,
		Duration:    s.Duration.Nanoseconds() / 1e3,
		Debug:       s.Debug,
		Source:      serviceToProto(s.Source),
		Destination: serviceToProto(s.Destination),
		Annotations: annotations,
	}
}

func (o *os) send(buf []*Span) {
	for _, s := range buf {
		pub := o.opts.Client.NewPublication(o.opts.Topic, toProto(s))
		o.opts.Client.Publish(context.TODO(), pub)
	}
}

func (o *os) run() {
	t := time.NewTicker(o.opts.BatchInterval)

	var buf []*Span

	for {
		select {
		case s := <-o.spans:
			buf = append(buf, s)
			if len(buf) >= o.opts.BatchSize {
				go o.send(buf)
				buf = nil
			}
		case <-t.C:
			// flush
			if len(buf) > 0 {
				go o.send(buf)
				buf = nil
			}
		case <-o.exit:
			t.Stop()
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

func (o *os) Collect(s *Span) error {
	select {
	case o.spans <- s:
	case <-time.After(o.opts.CollectTimeout):
		return errors.New("Timed out sending span")
	}
	return nil
}

func (o *os) NewSpan(s *Span) *Span {
	if s == nil {
		// completeley new trace
		return &Span{
			Id:        uuid.NewUUID().String(),
			TraceId:   uuid.NewUUID().String(),
			ParentId:  "0",
			Timestamp: time.Now(),
			Source:    o.opts.Service,
		}
	}

	// existing trace in theory
	cp := &Span{}
	*cp = *s

	if len(s.TraceId) == 0 {
		cp.TraceId = uuid.NewUUID().String()
	}
	if len(s.ParentId) == 0 {
		cp.ParentId = "0"
	}
	if len(s.Id) == 0 {
		cp.Id = uuid.NewUUID().String()
	}
	if s.Timestamp.IsZero() {
		cp.Timestamp = time.Now()
	}
	if s.Source == nil {
		cp.Source = o.opts.Service
	}

	return cp
}

func (o *os) String() string {
	return "os"
}
