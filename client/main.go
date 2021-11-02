package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/Mulac/encryptor/proto"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "localhost:50051", "The address to connect to")
	key  = flag.Int("key", 0, "The caesar cipher key to be used")
)

func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to %s: %v", *addr, err)
	}
	defer conn.Close()
	client := pb.NewEncryptorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// This is a very dumb and simple CLI
	// Cobra and Viper should be used for larger apps
	var res *pb.Message
	switch flag.Arg(0) {
	case "encrypt", "e":
		res, err = client.Encrypt(ctx, &pb.EncryptorRequest{Message: &pb.Message{Body: flag.Arg(1)}, Key: int32(*key)})
	case "decrypt", "d":
		res, err = client.Decrypt(ctx, &pb.EncryptorRequest{Message: &pb.Message{Body: flag.Arg(1)}, Key: int32(*key)})
	default:
		err = fmt.Errorf("unexpected input '%s'", flag.Arg(0))
	}
	if err != nil {
		log.Fatalf("==ERROR==\n%v\n\nUSAGE: go run client/main.go [options] <encrypt/decrypt> text\n", err)
	}

	fmt.Println(res.GetBody())
}
