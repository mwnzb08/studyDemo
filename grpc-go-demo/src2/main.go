package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/protoc"
	"log"
	"os"
)

const (
	address  = "localhost:9000"
	defaultName = "world"
)

func main() {
	mapss := make([]map[string]interface{}, 0)
	mapss = append(mapss, map[string]interface{}{"ss":"ss", "aa": "aa"})

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
	fmt.Println(r.Items)
	fmt.Println( r.Message)
	fmt.Println( r.String())
}