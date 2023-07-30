package main

import (
	pb "POC-rpc/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	PORT = ":3000"
)

type HelloServer struct {
	pb.GreetPackageServer
}

func (s *HelloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Greeting: "Hello" + req.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to setup tcp server %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetPackageServer(grpcServer, &HelloServer{})
	log.Printf("Server started on this %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start %v", err)
	}
}
