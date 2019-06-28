/*
@Time : 2019/4/23 15:27 
@Author : kenny zhu
@File : monitor.go
@Software: GoLand
@Others:
*/
package monitor

import (
	// "github.com/micro/go-micro/registry"
	"registry"
)

// Discovery builds on the registry for heart-beating and client metric
// all func here.
type Monitor interface {
	Close() error
    GetMonitorRegistry() registry.Registry
	PushHeartBeat(service *registry.Service, weights int)
}

type NewRegistry func(...registry.Option) registry.Registry
// var (
// 	// use et-cd instead if online.
// 	DefaultMonitor = NewMonitor(registry.NewRegistry)
// )

// this func open heart-beat
func NewMonitor(newRegistry NewRegistry, opts ...registry.Option) Monitor {
	return newOS(newRegistry, opts ...)
}

