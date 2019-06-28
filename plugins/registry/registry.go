/*
@Time : 2019/6/4 16:22
@Author : kenny zhu
@File : registry
@Software: GoLand
@Others:
*/
package registry

import (
	"errors"
)

type Option func(*Options)

type RegisterOption func(*RegisterOptions)

type WatchOption func(*WatchOptions)

var (
	// default use et-cdv3..
	DefaultRegistry = NewRegistry()

	// Not found error when GetService is called
	ErrNotFound = errors.New("not found")

	// Watcher stopped error when watcher is stopped
	ErrWatcherStopped = errors.New("watcher stopped")
)
