# Go OS [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/micro/go-os?status.svg)](https://godoc.org/github.com/micro/go-os) [![Travis CI](https://api.travis-ci.org/micro/go-os.svg?branch=master)](https://travis-ci.org/micro/go-os) [![Go Report Card](https://goreportcard.com/badge/micro/go-os)](https://goreportcard.com/report/github.com/micro/go-os)

Go OS is a client library for [Micro OS](https://github.com/micro/os)

NOTE: This is still a work in progress

## Usage

Each package is backed by a Micro OS service. Packages can be used independently or with go-micro wrappers.

## Features

Examples of usage can be found in [go-os/examples](https://github.com/micro/go-os/tree/master/examples)

Package     |   Built-in Plugin	|	Description
-------     |   --------	|	---------
[auth](https://godoc.org/github.com/micro/go-os/auth)	|	auth-srv	|   authentication and authorisation for users and services	
[config](https://godoc.org/github.com/micro/go-os/config)	|	config-srv	|   dynamic configuration which is namespaced and versioned
[db](https://godoc.org/github.com/micro/go-os/db)		|	db-srv		| distributed database abstraction
[discovery](https://godoc.org/github.com/micro/go-os/discovery)	|	discovery-srv	|   extends the go-micro registry to add heartbeating, etc
[event](https://godoc.org/github.com/micro/go-os/event)	|	event-srv	|	event publication, subscription and aggregation 
[kv](https://godoc.org/github.com/micro/go-os/kv)		|	distributed in-memory	|   simply key value layered on memcached, etcd, consul 
[log](https://godoc.org/github.com/micro/go-os/log)	|	file	|	structured logging to stdout, logstash, fluentd, pubsub
[monitor](https://godoc.org/github.com/micro/go-os/monitor)	|	monitor-srv	|   add custom healthchecks measured with distributed systems in mind
[metrics](https://godoc.org/github.com/micro/go-os/metrics)	|	telegraf	|   instrumentation and collation of counters
[router](https://godoc.org/github.com/micro/go-os/router)	|	router-srv	|	global circuit breaking, load balancing, A/B testing
[sync](https://godoc.org/github.com/micro/go-os/sync)	|	consul		|	distributed locking, leadership election, etc
[trace](https://godoc.org/github.com/micro/go-os/trace)	|	trace-srv	|	distributed tracing of request/response


