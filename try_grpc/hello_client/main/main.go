package main

import (
	pb "com.ledger.goproject/try_grpc/hello_server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "ledger1",
		"appKey": "123123",
	}, nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	var ops []grpc.DialOption

	ops = append(ops, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// 使用token进行传送
	ops = append(ops, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	conn, err := grpc.Dial("localhost:9090", ops...)
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()
	client := pb.NewGetMySonServiceClient(conn)
	req := &pb.LedgerRequest{
		Name: "ledger",
		Age:  18,
		Son: &pb.LedgerRequest_Son{
			Name: "son",
		},
	}
	son, err := client.GetSon(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", son)
}
