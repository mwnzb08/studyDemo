package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/protoc"
	"log"
	"net"
	"time"
)

type server struct {}
//
func main() {
	lis, err := net.Listen("tcp", ":8900")
	if err != nil {
		fmt.Println(err)
	}

	s := grpc.NewServer()
	protoc.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil  {
		fmt.Println(err)
	}

	ticker := time.NewTicker(time.Second * 1)
	for t:=range ticker.C{
		fmt.Println(t)
		
	}

}
// 注意继承方法的大小写SayHello No sayHello
func (s *server) SayHello(ctx context.Context, in *protoc.HelloRequest) (*protoc.HelloReply, error)  {
	log.Printf("receive : %v", in.Name)
	return &protoc.HelloReply{Message: "hello," + in.Name}, nil
}