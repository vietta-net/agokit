package gokit

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/metadata"
)




func TimezoneToContext(timezone string) grpctransport.ServerRequestFunc {
	if timezone != "" {
		ServerTimezone = timezone
	}
	return func(ctx context.Context, md metadata.MD) context.Context {
		// capital "Key" is illegal in HTTP/2.
		timezone, ok := md[TimezoneKey]
		tz := ServerTimezone
		if ok {
			tz = timezone[0]
		}

		ctx = context.WithValue(ctx, TimezoneKey, tz)
		return ctx
	}
}

func GetTimezoneFromContext(ctx context.Context) string {
	tz := ctx.Value(TimezoneKey)
	return tz.(string)
}

