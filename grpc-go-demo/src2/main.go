package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc/protoc"
	"log"
	"os"
)

const (
	address  = "localhost:8900"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := protoc.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &protoc.HelloRequest{Name: name, Value: "ss", Key: "saa", Num: 3})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}