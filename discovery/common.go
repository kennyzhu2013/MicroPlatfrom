/*
@Time : 2018/8/27 20:16 
@Author : kenny zhu
@File : common.go
@Software: GoLand
@Others:
*/
package discovery

import (
	"github.com/micro/go-micro/registry"
)

type   Heartbeat struct {
	id string// The ID of the heartbeat. Will generally be the Node ID.
	service *registry.Service // The service sending the heartbeat
	timestamp int64 // Unix time at which this was sent
	interval int64  // The interval at which this heartbeat is expected to be sent
	ttl int64  // The time to live for this heartbeat
}
