/*
@Time : 2019/4/19 16:28 
@Author : kenny zhu
@File : roundShardTripper.go
@Software: GoLand
@Others:
*/
package web

import (
	"errors"
	"fmt"
	"log/log"
	"net/http"
	"rcache"
	"registry"
	"time"
)

// st  selector.Strategy, with cache, act as selector.
type roundShardTripper struct {
	rt   http.RoundTripper
	opts Options
	rc rcache.Cache
}

func (r *roundShardTripper) newRCache() rcache.Cache {
	ropts := []rcache.Option{}
	if r.opts.Context != nil {
		if t, ok := r.opts.Context.Value("selector_ttl").(time.Duration); ok {
			ropts = append(ropts, rcache.WithTTL(t))
		}
	}
	return rcache.New(r.opts.Registry, ropts...)
}

// rewrite send request and receive response.
// req.URL.Host = er
func (r *roundShardTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// s, err := r.opts.Registry.GetService(req.URL.Host)
	s, err := r.rc.GetService(req.URL.Host)
	if err != nil {
		return nil, err
	}

	// get destination header.
	val := req.Header[r.opts.Destination]
	if len(val) < 1 {
		// no header, defer to client
		return r.roundTrip(s, req)
	}

	// get first server id.
	var serverId string
	for _, temp := range  val {
		// only get first..
		serverId = temp
		break
	}

	// split ip:port. with first one.
	// address := strings.Split(ipSets, ":")
	if len(serverId) < 1 {
		// valid serverId, defer to client
		return r.roundTrip(s, req)
	}
	// proxyIp := address[0]
	// proxyPort,_ := strconv.Atoi( address[1] )

	// Filter the nodes for serverTag marked by the server..
	var nodeResult *registry.Node
	for _, service := range s {
		for _, node := range service.Nodes {
			// if node.Address == proxyIp && node.Port == proxyPort {
			if node.Id == serverId {
				nodeResult = node
			}
		}
	}

	if nil == nodeResult {
		// not find any service
		log.Errorf("RoundTrip not find destination service:%v", serverId)
		return r.roundTrip(s, req)
	}

	// need retry ?...
	req.URL.Host = fmt.Sprintf("%s:%d", nodeResult.Address, nodeResult.Port)
	w, err := r.rt.RoundTrip(req)
	if err != nil {
		log.Error("RoundTrip failed request")
		return nil, errors.New("RoundTrip failed request")
	}

	// get destination header.
	// w.Header.Add(r.opts.Destination, req.URL.Host)
	w.Header.Add(r.opts.Destination, nodeResult.Id)
	log.Infof("RoundTrip done, header[%v] is set to: %v, host:%v!", r.opts.Destination, nodeResult.Id, req.URL.Host)
	return w, nil
}

// here use round-bin select
func (r *roundShardTripper)  roundTrip(s []*registry.Service, req *http.Request) (*http.Response, error) {
	// select the one with roundBinSelect
	next := r.opts.Selector(s)

	// rudimentary retry 3 times , may be the same one.
	for i := 0; i < 3; i++ {
		n, err := next()
		if err != nil {
			continue
		}
		if nil == n {
			log.Error("roundTrip failed not found any normal node")
			return nil, errors.New("roundTrip failed not found any normal node")
		}
		log.Infof("roundTrip found node with ip:%v, port:%v", n.Address, n.Port)
		req.URL.Host = fmt.Sprintf("%s:%d", n.Address, n.Port)

		// w, err := r.rt.RoundTrip(req)
		w, err := r.rt.RoundTrip(req)
		if err != nil {
			// de register
			// de register to fix?..
			r.deRegisterNode( s, n )
			log.Infof("roundTrip failed:%v!", err)
			continue
		}

		// if success, add media-server
		// w.Header.Add(r.opts.Destination, req.URL.Host)
		w.Header.Add(r.opts.Destination, n.Id)
		// w.Header[] = make([]string, 1, 1)
		// w.Header[r.opts.Destination][0] = req.URL.Host
		log.Infof("roundTrip success, header[%v] is set to: %v, host: %v!", r.opts.Destination, n.Id, req.URL.Host)
		return w, nil
	}

	log.Error("roundTrip failed request")
	return nil, errors.New("roundTrip failed request")
}

func (r *roundShardTripper) deRegisterNode(s []*registry.Service, node *registry.Node) {
	serviceNode := &registry.Service{
		Nodes: []*registry.Node{node},
	}

	for _, service := range s {
		for _, node := range service.Nodes {
			if node.Address == node.Address && node.Port == node.Port {
				serviceNode.Name = service.Name
				serviceNode.Version = service.Version

				// must update service node..
				_ = r.rc.Deregister( serviceNode )
				log.Infof("deRegisterNode id: %v", node.Id)
				break
			}
		}
	}
}