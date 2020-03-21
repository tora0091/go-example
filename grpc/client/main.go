package main

import (
	"context"
	"log"
	"os"
	"time"

	helloworld "go-example/grpc/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	name := "Masamune"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Greeting: %s\n", r.Message)

	r, err = c.SayHelloAgain(ctx, &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Greeting: %s\n", r.Message)
}
