package main

import (
	"com.ledger.goproject/try_grpc/grpcclient/sonos/v1/protocol"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	client := protocol.NewGreeterClient(conn)
	p := &protocol.HelloRequest{
		Name: "hh",
	}

	hello, err := client.SayHello(context.Background(), p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", hello)

}
