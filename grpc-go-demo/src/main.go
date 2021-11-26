package main

import (
	"context"
	"encoding/json"
	"fmt"
	structpb "github.com/golang/protobuf/ptypes/struct"
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

type Aaa struct {
	Aa string
	Ab string
	Ac string
}

// 注意继承方法的大小写SayHello No sayHello
func (s *server) SayHello(ctx context.Context, req *protoc.HelloRequest) (rsp *protoc.HelloReply, err error)  {
	itemsss := make(map[string][]map[string]interface{}, 0)
	itemsss["server"] = append(itemsss["server"], map[string]interface{}{
		"code": 1,
		"name": 2,
	})
	itemsss["server1"] = append(itemsss["server1"], map[string]interface{}{
		"code": 3,
		"name": 4,
	})
	marshal, err := json.Marshal(itemsss)
	if err != nil {
		return &protoc.HelloReply{}, nil
	}
	var a structpb.Struct
	err1 := json.Unmarshal(marshal, &a)
	if err1 != nil {
		return  &protoc.HelloReply{}, nil
	}
	if a.AsMap() != nil {
		fmt.Println(a.AsMap())
		rsp = &protoc.HelloReply{
			Items: &a,
		}
	}
	return
}
