package main

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	pb "github.com/tbocek/VSS-PROTO/grpc"
	"google.golang.org/grpc"
	"net"
)

type server struct{}

func (s *server) UserRPC(ctx context.Context, in *pb.User) (*pb.User, error) {
	in.Created = ptypes.TimestampNow()
	return in, nil
}

func main() {

	lis, err := net.Listen("tcp", "10.0.2.16:4000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
	s.Serve(lis)
}
