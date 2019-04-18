package main

import (
	"context"
	"time"
	"google.golang.org/grpc"
	pb "where/echolalia/.go/lives"
)

const (
	address	= "localhost:50051"
	message = "Hello there!"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEcholaliaClient(conn)
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Message: message})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response from server: %s", r.Message)
}
