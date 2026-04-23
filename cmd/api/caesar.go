package main

import (
	"context"
	"log"

	"github.com/basarardabaykal/cipherconverter-microservice/internal/cipher"
	"github.com/basarardabaykal/cipherconverter-microservice/internal/pb/common"
	"github.com/basarardabaykal/cipherconverter-microservice/internal/pb/symmetric"
)

func (s *server) EncryptCaesar(ctx context.Context, req *symmetric.CaesarRequest) (*common.CipherResponse, error) {
	log.Printf("Received EncryptCaesar with shift: %d", req.GetShift())

	c := cipher.NewCaesar(int(req.GetShift()))
	result := c.Encrypt(req.GetText())

	return &common.CipherResponse{Result: result}, nil
}

func (s *server) DecryptCaesar(ctx context.Context, req *symmetric.CaesarRequest) (*common.CipherResponse, error) {
	log.Printf("Received DecryptCaesar with shift: %d", req.GetShift())

	c := cipher.NewCaesar(int(req.GetShift()))
	result := c.Decrypt(req.GetText())

	return &common.CipherResponse{Result: result}, nil
}
