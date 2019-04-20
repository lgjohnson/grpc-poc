package main

import (
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/gjohnson/echolalia"
)

const port = ":50051"

type server struct {}

func (s *server) Echo(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply , error) {
	return &pb.HelloReply{Message: in.Message}, nil
}

func (s *server) Palindrome(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	word := in.Message
	rev := Reverse(word)
	return &pb.HelloReply{Message: word + rev[1:]}, nil
}
	
func Reverse(str string) string {
	if str != "" {
		return Reverse(str[1:]) + str[:1]
	}
	return ""   
}

func main() {
	listener, _ := net.Listen("tcp", port)
	s := grpc.NewServer()
	pb.RegisterEcholaliaServer(s, &server{})
	s.Serve(listener)
}
