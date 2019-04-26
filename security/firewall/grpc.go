package firewall

import (
	"context"

	"github.com/evalsocket/gophers-utility/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errInvalidRole = status.Errorf(codes.PermissionDenied, "invalid role")
)

// GrantGrpcAccessFor returns Status Unauthorized if
// Identity not set within request's context
// or user does not have required role
//
// 	The following grpc.ServerOption adds an interceptor for all unary RPCs.
//  To configure an interceptor for streaming RPCs, see:
// 	https://godoc.org/google.golang.org/grpc#StreamInterceptor
//
// opts := []grpc.ServerOption{
// 	grpc.UnaryInterceptor(GrantGrpcAccessFor("admin")),
// }
// s := grpc.NewServer(opts...)
// pb.RegisterGreeterServer(s, &server{})
func GrantGrpcAccessFor(role string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		i, ok := identity.FromContext(ctx)
		if ok {
			for _, userRole := range i.Roles {
				if userRole == role {
					return handler(ctx, req)
				}
			}
		}

		return nil, errInvalidRole
	}
}
