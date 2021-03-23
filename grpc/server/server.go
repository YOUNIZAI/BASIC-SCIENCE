// Package main implements a server for Greeter service.
package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/YOUNIZAI/BASIC-SCIENCE/pbfile"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

func main() {
	// 创建listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 创建server示例
	s := grpc.NewServer()
	// 注册服务
	pb.RegisterGreeterServer(s, &server{})
	// 启动服务端监听
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	if ctx.Err() == context.Canceled { //检查是否因为超过截止时间 RPC变为无效
		//s :=status.New("")
		//s.WithDetails()
		return nil,status.New(codes.Canceled, "Client cancelled, abandoning.").Err()
	}
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

