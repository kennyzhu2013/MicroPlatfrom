package monitor

import (
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	"time"

	"golang.org/x/net/context"
)

type clientWrapper struct {
	client.Client
	m Monitor
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	t := time.Now()
	err := c.Client.Call(ctx, req, rsp, opts...)
	c.m.RecordStat(req, time.Since(t), err)
	return err
}

func handlerWrapper(fn server.HandlerFunc, m Monitor) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		t := time.Now()
		err := fn(ctx, req, rsp)
		m.RecordStat(req, time.Since(t), err)
		return err
	}
}
