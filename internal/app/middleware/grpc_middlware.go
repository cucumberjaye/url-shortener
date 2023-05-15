package middleware

import (
	"context"

	"github.com/cucumberjaye/url-shortener/pkg/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// перехватчик аутентификации для grpc
func AuthenticationGRPC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	notAuth := []string{
		"/internal.app.pb.GetFullURL",
		"/internal.app.pb.Ping",
	}

	for _, val := range notAuth {
		if val == info.FullMethod {
			return handler(ctx, req)
		}
	}

	var authToken string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("authentication")
		if len(values) > 0 {
			authToken = values[0]
			if info.FullMethod == "/internal.app.pb.Authentication" {
				ctx = metadata.AppendToOutgoingContext(ctx, "authorization", authToken)
				return handler(ctx, req)
			}

			id, err := token.CheckToken(authToken)
			if err == nil {
				ctx = metadata.AppendToOutgoingContext(ctx, "user_id", id)
				return handler(ctx, req)
			}
		}
	}
	return nil, status.Error(codes.Unauthenticated, "unauthenticated")

}
