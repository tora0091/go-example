package main

import (
	"context"
	"go-example/grpc/proto"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
	SERVER_ADDR = "localhost:10049"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalln("error: parameter error")
	}
	code := os.Args[1]

	conn, err := grpc.Dial(SERVER_ADDR, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	foodClient := proto.NewFoodFactoryClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	foodReply, err := foodClient.GetFoodByID(ctx, &proto.FoodResponse{Code: code})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(foodReply)
}
