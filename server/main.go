package main

import (
	"log"
	"net"

	bp "grpc-server/service"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	bp.RegisterLongTimeRequestServiceServer(s, bp.NewDemoService())
	bp.RegisterRepeatedServiceServer(s, bp.NewRepeatedService())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
