package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"go-example/grpc/proto"
)

const (
	SERVER_ADDR = "localhost:10049"
)

type server struct{}

func (s *server) GetFoodByID(ctx context.Context, r *proto.FoodResponse) (*proto.FoodReply, error) {
	code := r.Code
	log.Printf("called function (GetFoodById), code: %s\n", code)

	foods := []*proto.FoodReply{
		&proto.FoodReply{Code: "A001", Name: "Big Mountain Pizza", Price: 399, Other: "Many Big Cheese", Drink: map[string]string{"D001": "Pepsi", "D002": "Fanta", "D003": "Milk"}},
		&proto.FoodReply{Code: "A002", Name: "No Meet Humburger", Price: 1299, Other: "it is natural. very good taste", Drink: map[string]string{"D004": "Green tea", "D009": "Water"}},
		&proto.FoodReply{Code: "A003", Name: "Deep-fried chicken and rice", Price: 649, Other: "i love it", Drink: map[string]string{"D001": "Pepsi", "D005": "Diet Cola"}},
		&proto.FoodReply{Code: "X001", Name: "Sushi and Green tea", Price: 599, Other: "it is japanese", Drink: map[string]string{"D006": "Green tea japanese version"}},
		&proto.FoodReply{Code: "Z001", Name: "Morning set low price", Price: 99, Other: "Good Morning!!", Drink: map[string]string{"D007": "cafe", "D008": "morning cafe", "D009": "latte"}},
	}

	var target *proto.FoodReply
	target = nil
	for _, food := range foods {
		if food.Code == code {
			target = food
			break
		}
	}
	if target == nil {
		return nil, fmt.Errorf("Food Not Found: code %s", code)
	}
	return target, nil
}

func main() {
	listen, err := net.Listen("tcp", SERVER_ADDR)
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	proto.RegisterFoodFactoryServer(s, &server{})

	log.Printf("open food factory: %s\n", SERVER_ADDR)
	log.Fatalln(s.Serve(listen))
}
