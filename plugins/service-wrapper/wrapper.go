/*
@Time : 2019/5/13 10:36 
@Author : kenny zhu
@File : wrapper.go
@Software: GoLand
@Others:
*/
package service_wrapper

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

// Service is a web service with service discovery built in
type Service interface {
	Client() *http.Client
	Options() Options
	Run() error
}

type Option func(o *Options)
type Runner func(addr ...string) (err error)

var (
	// For serving
	DefaultName    = "gin-web"
	DefaultVersion = "latest"
	DefaultId      = uuid.New().String()
	DefaultAddress = ":0"

	// for registration
	DefaultRegisterTTL      = time.Minute
	DefaultRegisterInterval = time.Second * 30
)

// NewService returns a new web.ServiceWrapper, for future use service for extend.
func NewService(opts ...Option) Service {
	return newService(opts...)
}
