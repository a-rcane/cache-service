package main

import (
	"context"
	"fmt"
	"main/proto"
	"net"
	"google.golang.org/grpc"
	"errors"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp",":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterBasicServiceServer(srv, &server{})
	fmt.Println("Prepare to serve")
	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *server) Set(ctx context.Context, in *setReq.key, in *setReq.value) error {
    return errors.New("not implemented yet. Sanchit will implement me")
}

func (s *server) Get(ctx context.Context, in *getReq.key) (*getRes.value, error) {
    return []bytes("Sanchit will implement me"), nil
}