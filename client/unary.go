package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Sanskarzz/grpcbasic/proto"
)

func callSayHellp(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}
