package jwt

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (m *Middleware) UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	header, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.New(codes.Unauthenticated, "no auth provided").Err()
	}

	tokens := header.Get("jwt")
	if len(tokens) < 1 {
		return nil, status.New(codes.Unauthenticated, "no auth provided").Err()
	}
	tokenString := tokens[0]

	token, err := m.GetToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	ctx = ContextWithToken(ctx, token)
	fmt.Println("* gRPC SERVER middleware validated and set token")
	return handler(ctx, req)
}
