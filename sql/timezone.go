package sql

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/metadata"
	"context"
)

func TimeZoneToContext(timezone *string) grpctransport.ServerRequestFunc {
	return func(ctx context.Context, md metadata.MD) context.Context {
		// capital "Key" is illegal in HTTP/2.
		timezone, ok := md["timezone"]
		if !ok {
			return ctx
		}
		ctx = context.WithValue(ctx, "timezone", timezone)
		return ctx
	}
}