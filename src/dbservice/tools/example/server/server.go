// for proteus go-grpc server..

package server

import (
	example2 "github.com/kennyzhu/go-os/src/dbservice/tools/example"
	"net"

	"google.golang.org/grpc"
)

// NewServer returns a new server serving in the given address
func NewServer(addr string) (*grpc.Server, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	grpcServer := grpc.NewServer()
	example2.RegisterExampleServiceServer(grpcServer, example2.NewExampleServiceServer())
	go grpcServer.Serve(lis)
	return grpcServer, nil
}
