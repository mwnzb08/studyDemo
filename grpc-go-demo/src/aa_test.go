package main

import (
	"context"
	"grpc/protoc"
	"testing"
)

func BenchmarkFunc(t *testing.B) {
	a := new(server)
	for i :=0;i<t.N;i++ {
		a.SayHello(context.Background(), &protoc.HelloRequest{Name: "name", Value: "ss", Key: "saa", Num: 3})
	}

}