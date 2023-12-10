package service

import (
	"com.ledger.goproject/try_grpc/grpcservice/sonos/v1/protocol"
	"context"
)

type HelloGrpc struct {
	protocol.UnimplementedGreeterServer
}

func (h HelloGrpc) Return(ctx context.Context, info *protocol.ResultInfo) (*protocol.ResultResponse, error) {
	p := &protocol.ResultResponse{
		Result: &protocol.ResultInfo{
			Code:    "9000",
			Message: "message",
		},
	}
	return p, nil
}

func NewHelloGrpc() *HelloGrpc {
	return &HelloGrpc{}
}

func (h HelloGrpc) SayHello(ctx context.Context, request *protocol.HelloRequest) (*protocol.HelloResponse, error) {
	p := &protocol.HelloResponse{
		Greeting: request.Name + "这个是你的名字!!!",
	}
	return p, nil
}
func (h HelloGrpc) SayBye(ctx context.Context, request *protocol.ByeRequest) (*protocol.ByeResponse, error) {

	return &protocol.ByeResponse{
		Name:      request.Name,
		Age:       0,
		IsStudent: false,
		Hobbies: []string{
			"吃饭", "睡觉", "大豆都",
		},
		Address: &protocol.Address{
			Street:  "",
			City:    "",
			Country: "chind",
		},
	}, nil
}
