package main

import(
	"context"
	"fmt"
	"main/proto"
	"net"
	"google.golang.org/grpc"
	"errors"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := NewMakeServiceClient(conn) 

	func (s *server) Set(ctx context.Context, in *setReq.key, in *setReq.value) error {
	    return errors.New("not implemented yet. Sanchit will implement me")
	}

	func (s *server) Get(ctx context.Context, in *getReq.key) (*getRes.value, error) {
	    return []bytes("Sanchit will implement me"), nil
	}

}
