package middleware

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TokenMiddleware(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp any, err error) {
	err = Auth(ctx)
	if err != nil {
		return
	}
	return handler(ctx, req)
}

func Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}
	var user string
	var password string
	if val, ok := md["user"]; ok {
		user = val[0]
	}
	if val, ok := md["password"]; ok {
		password = val[0]
	}
	if user != "admin" || password != "admin" {
		return status.Errorf(codes.Unauthenticated, "客户端请求的token不合法")
	}
	return nil
}
