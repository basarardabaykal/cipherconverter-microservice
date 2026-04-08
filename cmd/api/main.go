package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/basarardabaykal/cipherconverter-microservice/internal/pb/symmetric"
)

type symServer struct {
	symmetric.UnimplementedCipherServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	symmetric.RegisterCipherServiceServer(grpcServer, &symServer{})

	log.Printf("gRPC Cipher microservice is running and listening on %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
