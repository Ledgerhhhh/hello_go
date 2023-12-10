package main

import (
	"com.ledger.goproject/try_grpc/grpcservice/middleware"
	"com.ledger.goproject/try_grpc/grpcservice/sonos/v1/protocol"
	"com.ledger.goproject/try_grpc/grpcservice/sonos/v1/service"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	// 创建网络监听器，监听指定地址（0.0.0.0 表示监听所有网络接口）和端口（8888）。
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		// 如果监听失败，输出错误信息并终止程序。
		fmt.Println(err)
		return
	}
	// rpcServer 是一个 gRPC 服务器实例，用于处理 gRPC 请求。
	rpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.TokenMiddleware))

	// helloGrpc 是你定义的服务实例，实现了 GreeterServer 接口中的具体服务端方法。
	helloGrpc := service.NewHelloGrpc()

	// 将 helloGrpc 注册到 gRPC 服务器 rpcServer 中，以便处理客户端的 gRPC 请求。
	protocol.RegisterGreeterServer(rpcServer, helloGrpc)

	// 启动 gRPC 服务器，开始监听客户端的连接，并处理 gRPC 请求。
	rpcServer.Serve(listen)

}
