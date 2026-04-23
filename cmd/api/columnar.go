package main

import (
	"context"
	"log"

	"github.com/basarardabaykal/cipherconverter-microservice/internal/cipher"
	"github.com/basarardabaykal/cipherconverter-microservice/internal/pb/common"
	"github.com/basarardabaykal/cipherconverter-microservice/internal/pb/symmetric"
)

func (s *server) EncryptColumnar(ctx context.Context, req *symmetric.ColumnarRequest) (*common.CipherResponse, error) {
	log.Printf("Received EncryptColumnar with %d columns", req.GetColumns())

	c := cipher.NewUnkeyedColumnarTransposition(int(req.GetColumns()))
	result := c.Encrypt(req.GetText())

	return &common.CipherResponse{Result: result}, nil
}

func (s *server) DecryptColumnar(ctx context.Context, req *symmetric.ColumnarRequest) (*common.CipherResponse, error) {
	log.Printf("Received DecryptColumnar with %d columns", req.GetColumns())

	c := cipher.NewUnkeyedColumnarTransposition(int(req.GetColumns()))
	result := c.Decrypt(req.GetText())

	return &common.CipherResponse{Result: result}, nil
}
