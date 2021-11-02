package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/Mulac/encryptor/encryptor"
	pb "github.com/Mulac/encryptor/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var (
	addr  = flag.String("addr", ":50051", "The address to listen to")
	crypt = flag.String("encryptor", "caesar", "The encryptor to use")
)

type encryptorServer struct {
	pb.UnimplementedEncryptorServer
	encryptor.Encryptor
}

func (s *encryptorServer) Encrypt(ctx context.Context, req *pb.EncryptorRequest) (res *pb.Message, err error) {
	// primative logging
	log.Printf("INFO|server.Encrypt|Received: %v", req)
	defer func() {
		log.Printf("INFO|server.Encrypt|Finished: %v|returned %s, %v", req, res, err)
	}()

	encrypted, err := s.Encryptor.Encrypt(req.Message.Body, encryptor.Key(int(req.Key)))
	if errors.Is(err, encryptor.ErrKey) {
		// Key provided was invalid
		return nil, status.Error(3, fmt.Sprintf("failed to encrypt %s: alphabet rotation factor is 0", req))
	}
	if err != nil {
		// Internal server error
		log.Println(err)
		return nil, status.Error(13, fmt.Sprintf("failed to encrypt %s", req))
	}

	return &pb.Message{Body: encrypted}, nil
}

func (s *encryptorServer) Decrypt(ctx context.Context, req *pb.EncryptorRequest) (res *pb.Message, err error) {
	// primative logging
	log.Printf("INFO|server.Decrypt|Received: %v", req)
	defer func() {
		log.Printf("INFO|server.Decrypt|Finished: %v|returned %s, %v", req, res, err)
	}()

	decrypted, err := s.Encryptor.Decrypt(req.Message.Body, encryptor.Key(int(req.Key)))
	if errors.Is(err, encryptor.ErrKey) {
		// Key provided was invalid
		return nil, status.Error(3, fmt.Sprintf("failed to decrypt %s: alphabet rotation factor is 0", req))
	}
	if err != nil {
		// Internal server error
		log.Println(err)
		return nil, status.Error(13, fmt.Sprintf("failed to decrypt %s", req))
	}
	return &pb.Message{Body: decrypted}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen at %s: %v", *addr, err)
	}

	// TODO(calum): add logging middleware and TLS support
	s := grpc.NewServer()

	// TODO(calum): get encryptor type from config
	e, err := encryptor.NewEncryptor(encryptor.EncryptorType(*crypt))
	if err != nil {
		log.Fatalf("failed to create caesar encryptor: %v", err)
	}

	pb.RegisterEncryptorServer(s, &encryptorServer{Encryptor: e})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
