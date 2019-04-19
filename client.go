package main

import (
	"log"
	"context"
	"time"
	"google.golang.org/grpc"
	pb "github.com/gjohnson/echolalia"
)

const (
	address	= "localhost:50051"
	message = "race"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEcholaliaClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Palindrome(ctx, &pb.HelloRequest{Message: message})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response from server: %s", r.Message)
}
