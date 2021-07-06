package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"grpc/protoc"
	"net"
	"time"
)

type server struct {}
//
func main() {
	lis, err := net.Listen("tcp", ":9000")
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
	mapp := make(map[string]interface{},0)
	aaa,_ := json.Marshal(*in)
	json.Unmarshal(aaa, &mapp)
	//fmt.Println(mapp)
	//request := strings.Fields(in.String())
	//for _, ss := range request {
	//	dd := strings.Split(ss,":")
	//	mapp[dd[0]] = dd[1]
	//}
	//fmt.Println(mapp)
	return &protoc.HelloReply{Message: "hello," + in.Name}, nil
}
