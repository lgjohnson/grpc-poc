package main

import (
	"context"
	"net"
	"google.golang.org/grpc"
	pb "where/proto/code/is"
)

const port = ":50051"

type server struct {}

func (s *server) Echo(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply , error) {
	return &pb.HelloReply{Message: in.Message}, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	s := grpc.NewServer()
	pb.RegisterEcholaliaServer(s, &server{})
	s.Serve(listener)
}
