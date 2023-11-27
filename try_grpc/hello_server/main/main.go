package main

import (
	pb "com.ledger.goproject/try_grpc/hello_server/proto"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
)

type serve struct {
	pb.UnimplementedHelloServiceServer
	pb.UnimplementedGetMySonServiceServer
}

// MyInterceptor 是自定义的拦截器
type MyInterceptor struct{}

// UnaryInterceptor 实现了一元拦截器的接口
func (i *MyInterceptor) UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("FromIncomingContext err")
		return nil, errors.New("err auth")
	}
	appid := md.Get("appid")
	appkey := md.Get("appkey")
	if appid[0] != "ledger" || appkey[0] != "123123" {
		return nil, errors.New("err auth")
	}

	resp, err = handler(ctx, req)

	return resp, err
}

// GetSon StreamInterceptor 实现了流拦截器的接口
//
//	func (i *MyInterceptor) StreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
//		fmt.Printf("StreamInterceptor: Received request for %s\n", info.FullMethod)
//		err := handler(srv, ss)
//		fmt.Printf("StreamInterceptor: Sent response for %s\n", info.FullMethod)
//		return err
//	}
func (s serve) GetSon(ctx context.Context, request *pb.LedgerRequest) (*pb.Son, error) {
	log.Printf("%+v", request)

	return &pb.Son{
		Name: request.Son.Name,
	}, nil
}

func (s serve) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("%+v", request)

	return &pb.HelloResponse{
		Greet:       "hello",
		Age:         18,
		ResponseMsg: "hhh",
	}, nil
}

func main() {
	// 创建证书
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println(err)
	}
	// 注册拦截器
	grpcServer := grpc.NewServer(
		// 一元拦截器
		grpc.UnaryInterceptor((&MyInterceptor{}).UnaryInterceptor),
		// 流拦截器
		//grpc.StreamInterceptor((&MyInterceptor{}).StreamInterceptor),
	)
	pb.RegisterHelloServiceServer(grpcServer, &serve{})
	pb.RegisterGetMySonServiceServer(grpcServer, &serve{})
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Println(err)
	}
}
