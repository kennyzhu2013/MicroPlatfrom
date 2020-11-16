/*
@Time : 2019/6/21 10:22
@Author : kenny zhu
@File : options
@Software: GoLand
@Others:
*/
package rcache

import "time"

type Options struct {
	// TTL is the cache TTL
	TTL time.Duration
}

// WithTTL sets the cache TTL
func WithTTL(t time.Duration) Option {
	return func(o *Options) {
		o.TTL = t
	}
}