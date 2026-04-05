package main

import (
	"context"
	"log"

	"github.com/basarardabaykal/cipherconverter-microservice/internal/cipher"
	"github.com/basarardabaykal/cipherconverter-microservice/internal/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) EncryptOTP(ctx context.Context, req *pb.OTPRequest) (*pb.CipherResponse, error) {
	if len(req.GetText()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "text to encrypt cannot be empty")
	}
	if len(req.GetKey()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "OTP key must be provided")
	}

	log.Printf("Received EncryptOTP (text length: %d, key length: %d)", len(req.GetText()), len(req.GetKey()))

	c := cipher.NewOTP(req.GetKey())
	result := c.Encrypt(req.GetText())

	return &pb.CipherResponse{Result: result}, nil
}

func (s *server) DecryptOTP(ctx context.Context, req *pb.OTPRequest) (*pb.CipherResponse, error) {
	if len(req.GetText()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "text to decrypt cannot be empty")
	}
	if len(req.GetKey()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "OTP key must be provided")
	}

	log.Printf("Received DecryptOTP (text length: %d, key length: %d)", len(req.GetText()), len(req.GetKey()))

	c := cipher.NewOTP(req.GetKey())
	result := c.Decrypt(req.GetText())

	return &pb.CipherResponse{Result: result}, nil
}
