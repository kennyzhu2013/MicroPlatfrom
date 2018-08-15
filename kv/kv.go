// Package kv is an interface for key-value storage.
package kv

import (
	"errors"
	"time"
)

var (
	ErrNotFound = errors.New("not found")
)

type KV interface {
	Close() error
	Get(key string) (*Item, error)
	Del(key string) error
	Put(item *Item) error
	String() string
}

type Item struct {
	Key        string
	Value      []byte
	Expiration time.Duration
}

type Option func(o *Options)

func NewKV(opts ...Option) KV {
	return newOS(opts...)
}
