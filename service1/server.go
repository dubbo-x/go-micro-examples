package main

import (
	"go-micro-examples/service1/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type greeter struct{}

func (s *greeter) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: request.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &greeter{})
	s.Serve(listen)
}
