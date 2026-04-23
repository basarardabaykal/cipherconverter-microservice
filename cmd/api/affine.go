package main

import (
	"context"
	"log"

	"github.com/basarardabaykal/cipherconverter-microservice/internal/cipher"
	"github.com/basarardabaykal/cipherconverter-microservice/internal/pb/common"
	"github.com/basarardabaykal/cipherconverter-microservice/internal/pb/symmetric"
)

func (s *server) EncryptAffine(ctx context.Context, req *symmetric.AffineRequest) (*common.CipherResponse, error) {
	log.Printf("Received EncryptAffine with a: %d b: %d", req.GetA(), req.GetB())

	c := cipher.NewAffine(int(req.GetA()), int(req.GetB()))
	result := c.Encrypt(req.GetText())

	return &common.CipherResponse{Result: result}, nil
}

func (s *server) DecryptAffine(ctx context.Context, req *symmetric.AffineRequest) (*common.CipherResponse, error) {
	log.Printf("Received DecryptAffine with a: %d b: %d", req.GetA(), req.GetB())

	c := cipher.NewAffine(int(req.GetA()), int(req.GetB()))
	result := c.Decrypt(req.GetText())

	return &common.CipherResponse{Result: result}, nil
}
