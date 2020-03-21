package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	helloworld "go-example/grpc/proto"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, r *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello, " + r.Name + "!"}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, r *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello Again, " + r.Name + "!"}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
