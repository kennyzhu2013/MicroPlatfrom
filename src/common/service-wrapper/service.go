/*
@Time : 2019/5/13 10:46 
@Author : kenny zhu
@File : service.go
@Software: GoLand
@Others:
*/
package service_wrapper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"log/log"
	"net/http"
	"os"
	"os/signal"
	"registry"
	"sync"
	"syscall"
	"time"
	"web"

	"strconv"
	. "strings"
	maddr "util/addr"
)

// gin based..
type service struct {
	opts Options

	// mux *http.ServeMux
	mux *http.ServeMux
	srv *registry.Service

	sync.Mutex
	running bool
}

func newService(opts ...Option) Service {
	options := newOptions(opts...)
	s := &service{
		opts: options,
		mux:  http.NewServeMux(),
	}

	if s.opts.ServiceInfo != nil {
		s.srv = s.opts.ServiceInfo
	} else {
		s.srv = s.genSrv()
	}
	return s
}

// gin service init..
func (s *service) genSrv() *registry.Service {
	// default host:port
	parts := Split(s.opts.Address, ":")

	// support many ip-s
	host := Join(parts[:len(parts)-1], ":")
	port, _ := strconv.Atoi(parts[len(parts)-1])

	// check the advertise address first
	// if it exists then use it, otherwise
	// use the address
	if len(s.opts.Advertise) > 0 {
		parts = Split(s.opts.Advertise, ":")

		// we have host:port
		if len(parts) > 1 {
			// set the host
			host = Join(parts[:len(parts)-1], ":")

			// get the port
			if aport, _ := strconv.Atoi(parts[len(parts)-1]); aport > 0 {
				port = aport
			}
		} else {
			host = parts[0]
		}
	}

	addr, err := maddr.Extract(host)
	if err != nil {
		// best effort localhost, default
		addr = "127.0.0.1"
	}

	// format last address
	s.opts.Address = addr + ":" + strconv.Itoa(port)

	return &registry.Service{
		Name: s.opts.Name,
		Metadata: map[string]string{
			"serverDescription": s.opts.Description,  // server desc.
		},
		Nodes: []*registry.Node{
			{
				Id:       "go.micro.api.media-proxy-",
				Address:  addr,
				Port:     port,
				Metadata: s.opts.Metadata,
			},
		},
		Version: s.opts.Version,
	}
}

// heart beat run here, use monitor instead.
func (s *service) run(exit chan bool) {
	if s.opts.RegisterInterval <= time.Duration(0) {
		// use HeartBeatCheck if no
		return
	}

	t := time.NewTicker(s.opts.RegisterInterval)
	for {
		select {
		case <-t.C:
			// just retry circle if failed
			_ = s.register()
		case <-exit:
			t.Stop()
			return
		}
	}
}

func (s *service) register() error {
	if s.srv == nil {
		return nil
	}
	// default to service registry
	r := s.opts.Registry
	if s.opts.Registry == nil {
		return errors.New("Registry is empty. ")
	}
	return r.Register(s.srv, registry.RegisterTTL(s.opts.RegisterTTL))
}

func (s *service) deregister() error {
	if s.srv == nil {
		return nil
	}
	// default to service registry
	r := s.opts.Registry
	if s.opts.Registry == nil {
		return errors.New("Registry is empty. ")
	}
	return r.Deregister(s.srv)
}

// start server..
func (s *service) start() error {
	s.Lock()
	defer s.Unlock()

	if s.running {
		return nil
	}

	if s.opts.Engine == nil {
		gin.SetMode(gin.ReleaseMode)
		s.opts.Engine = gin.Default()
	}
	if s.opts.Secure  {
		s.opts.Engine.Use( s.LoadTLS() )
		go s.opts.Engine.RunTLS(s.opts.Address, s.opts.TLSConfig.CertFile, s.opts.TLSConfig.KeyFile)
	} else  {
		go s.opts.Engine.Run(s.opts.Address)
	}

	// s.exit = make(chan chan error, 1)
	s.running = true
	// go func() {
	// 	ch := <-s.exit
	//
	// 	ch <-
	// }()

	log.Infof("Listening on %v\n", s.opts.Address)
	return nil
}

func (s *service) stop()  {
	s.Lock()
	defer s.Unlock()
	if !s.running {
		return
	}

	s.running = false

	log.Info("Stopping")
	return
}

// load tls
func (s *service) LoadTLS() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     s.opts.Address,
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}

// for http-client, must add own rt.
func (s *service) Client(opts ...web.Option) *http.Client {
	// use random selector and  replace http.NewRoundTripper with web tripper
	rt := web.NewRoundShardTripper( opts... )

	return &http.Client{
		Transport: rt,
	}
}

// main run
func (s *service) Run() error {
	if err := s.start(); err != nil {
		return err
	}

	if err := s.register(); err != nil {
		return err
	}

	// start reg monitor loop
	ex := make(chan bool)
	go s.run(ex)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	select {
	// wait on kill signal
	case sig := <-ch:
		log.Info("Received signal %s\n", sig)
	// wait on context cancel
	case <-s.opts.Context.Done():
		log.Info("Received context shutdown")
	}

	// exit reg loop
	close(ex)
	if err := s.deregister(); err != nil {
		return err
	}

	s.stop()
	return nil
}

// Options returns the options for the given service
func (s *service) Options() Options {
	return s.opts
}
