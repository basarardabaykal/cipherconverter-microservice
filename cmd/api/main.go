package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/basarardabaykal/cipherconverter-microservice/internal/pb/symmetric"
)

type server struct {
	symmetric.UnimplementedCipherServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	symmetric.RegisterCipherServiceServer(s, &server{})

	log.Printf("gRPC Cipher microservice is running and listening on %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
