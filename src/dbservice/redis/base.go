/*
@Time : 2018/9/4 12:41 
@Author : kenny zhu
@File : base.go
@Software: GoLand
@Others: https://godoc.org/github.com/gomodule/redigo/redis#pkg-examples
*/
package redis

import (
	cluster "github.com/chasex/redis-go-cluster"
	"time"
)

// any call here..
type  RedisService interface {
	GetInstance() *cluster.Cluster
	GetBatch() *cluster.Batch
	Do(commandName string, args ...interface{}) ( interface{},  error)
}

// NewRedisService returns the default redis service.
func newRedisService(clusterIn *cluster.Cluster) RedisService {
	return &redisService{ defaultCluster:clusterIn, }
}

type redisService struct {
	// repo repositories.MovieRepository
	defaultCluster *cluster.Cluster
}
var RedisServe RedisService


func Init( addresses []string )  {
	defaultCluster,_ := cluster.NewCluster(
		&cluster.Options{
		StartNodes: addresses, // []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002"},
		ConnTimeout: 50 * time.Millisecond,
		ReadTimeout: 50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
		KeepAlive: 16,
		AliveTime: 60 * time.Second,
	})
	RedisServe = newRedisService(defaultCluster)
}

func(s *redisService) GetInstance() *cluster.Cluster {
	return s.defaultCluster
}

func(s *redisService) GetBatch() *cluster.Batch {
	return s.defaultCluster.NewBatch()
}

// commands: https://redis.io/commands.
func(s *redisService) Do(commandName string, args ...interface{}) ( interface{},  error) {
	return s.defaultCluster.Do(commandName, args...)
}

// all other operations : https://godoc.org/github.com/gomodule/redigo/redis#pkg-examples..