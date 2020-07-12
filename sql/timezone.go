package sql

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/metadata"
)

const TimezoneKey = "timezone"

func TimezoneToContext(timezone *string) grpctransport.ServerRequestFunc {
	return func(ctx context.Context, md metadata.MD) context.Context {
		// capital "Key" is illegal in HTTP/2.
		timezone, ok := md[TimezoneKey]
		if !ok {
			return ctx
		}
		ctx = context.WithValue(ctx, TimezoneKey, timezone)
		return ctx
	}
}

func GetTimezoneFromContext(ctx context.Context) string {
	tz := ctx.Value(TimezoneKey)
	return tz.(string)
}
